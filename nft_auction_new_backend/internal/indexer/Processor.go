package indexer

import (
	"backend/internal/db"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"log"
	"math/big"
	"strings"
	"time"
)

// 事件处理器
type Processor struct {
	db       *db.MySQL
	gorm     *gorm.DB
	chainID  uint64
	contract string
	abi      abi.ABI
}

func NewProcessor(db *db.MySQL, chainID uint64, contractAddress string) *Processor {
	return &Processor{
		db:       db,
		gorm:     db.DB,
		chainID:  chainID,
		contract: strings.ToLower(contractAddress),
		abi:      mustParseABI(),
	}
}

func (p *Processor) handleLog(ctx context.Context, lg types.Log) error {
	log.Printf(
		"HandleLog: tx=%s topic0=%s block=%d",
		lg.TxHash.Hex(),
		lg.Topics[0].Hex(),
		lg.BlockNumber,
	)

	if len(lg.Topics) == 0 {
		return nil
	}

	event, err := p.abi.EventByID(lg.Topics[0])
	if err != nil {
		return nil
	}

	switch event.Name {
	case "CreateAuction":
		return p.handleAuctionCreated(&lg)
	case "BidAuction":
		return p.handleBidPlaced(&lg)
	case "AuctionEnded":
		return p.handleAuctionEnded(&lg)
	case "AuctionCancelled":
		return p.handleAuctionCancelled(&lg)
	default:
		return nil
	}
}

// AuctionCreated event handle
func (p *Processor) handleAuctionCreated(lg *types.Log) error {
	ctx := context.Background()

	//1.非indexed 数据解析
	var data struct {
		TokenId     *big.Int
		EndTime     *big.Int
		MinBidUsd18 *big.Int
	}

	if err := p.abi.UnpackIntoInterface(&data, "CreateAuction", lg.Data); err != nil {
		return fmt.Errorf("unpack CreateAuction event failed: %w", err)
	}

	// 2. topics：indexed 参数
	if len(lg.Topics) < 4 {
		return fmt.Errorf("invalid CreateAuction topics len=%d", len(lg.Topics))
	}

	auctionId := new(big.Int).SetBytes(lg.Topics[1].Bytes()).Uint64()
	seller := common.BytesToAddress(lg.Topics[2].Bytes())
	nft := common.BytesToAddress(lg.Topics[3].Bytes())

	// 3. 防御性校验
	if data.TokenId == nil || data.EndTime == nil || data.MinBidUsd18 == nil {
		return fmt.Errorf("AuctionCreated data has nil field: %+v", data)
	}

	// 4. 构造 auction 对象
	auction := db.Auction{
		ChainID:         p.chainID,
		ContractAddress: p.contract,
		AuctionID:       auctionId,
		NFTContract:     nft.Hex(),
		TokenID:         data.TokenId.String(),
		Seller:          seller.Hex(),
		StartTime:       time.Unix(int64(lg.BlockNumber), 0), // 临时用 blockNumber
		EndTime:         time.Unix(data.EndTime.Int64(), 0),
		MinBidUsd18:     data.MinBidUsd18.String(),
		Status:          "OPEN",
		CreatedTxHash:   lg.TxHash.Hex(),
		CreatedBlock:    lg.BlockNumber,
	}

	// 5. 写数据库
	return p.db.UpsertAuctionWithLog(ctx, &auction)
}

// BidPlaced event handle 买家出价
func (p *Processor) handleBidPlaced(lg *types.Log) error {
	//1. 解析非 indexed data
	var data struct {
		BidToken  common.Address
		BidAmount *big.Int
		BidUsd18  *big.Int
	}

	if err := p.abi.UnpackIntoInterface(&data, "BidAuction", lg.Data); err != nil {
		return fmt.Errorf("unpack BidAuction failed: %w", err)
	}

	// 2. 解 indexed topics
	if len(lg.Topics) < 3 {
		return fmt.Errorf("invalid BidPlaced topics len=%d", len(lg.Topics))
	}

	auctionId := new(big.Int).SetBytes(lg.Topics[1].Bytes()).Uint64()
	bidder := common.BytesToAddress(lg.Topics[2].Bytes())

	// 3. 参数校验
	if data.BidAmount == nil || data.BidUsd18 == nil {
		return fmt.Errorf("BidPlaced data has nil field: %+v", data)
	}

	// 4.写库bid 构造
	bid := &db.Bid{
		ChainID:         p.chainID,
		ContractAddress: p.contract,
		AuctionID:       auctionId,
		Bidder:          strings.ToLower(bidder.Hex()),
		BidToken:        strings.ToLower(data.BidToken.Hex()),
		BidAmount:       data.BidAmount.String(),
		BidUsd18:        data.BidUsd18.String(),
		TxHash:          lg.TxHash.Hex(),
		BlockNumber:     lg.BlockNumber,
		BidTime:         time.Now().UTC(),
	}

	//5. 写入数据库
	if err := p.db.InsertBid(context.Background(), bid); err != nil {
		return fmt.Errorf("insert bid failed: %w", err)
	}

	//6.调用Repo 更新最高出价信息
	return p.db.UpdateAuctionOnBid(
		context.Background(),
		p.chainID,
		p.contract,
		auctionId,
		strings.ToLower(data.BidToken.Hex()),
		data.BidAmount.String(),
		data.BidUsd18.String(),
		strings.ToLower(bidder.Hex()),
		lg.BlockNumber)
}

// AuctionEnded event handle 拍卖结束事件处理
func (p *Processor) handleAuctionEnded(lg *types.Log) error {

	ctx := context.Background()

	//1.解析非indexed data
	var da struct {
		Winner    common.Address
		PayToken  common.Address
		PayAmount *big.Int
		PayUsd18  *big.Int
	}

	if err := p.abi.UnpackIntoInterface(&da, "AuctionEnded", lg.Data); err != nil {
		return fmt.Errorf("unpack AuctionEnded failed: %w", err)
	}

	//2. 解析 indexed topics
	if len(lg.Topics) < 2 {
		return fmt.Errorf("invalid AuctionEnded topics len=%d", len(lg.Topics))
	}

	auctionId := new(big.Int).SetBytes(lg.Topics[1].Bytes()).Uint64()

	//3. 调用repo 更新拍卖状态 "ENDED"
	if err := p.db.MarkAuctionEnded(
		ctx,
		p.chainID,
		p.contract,
		auctionId,
		strings.ToLower(da.Winner.Hex()),
		lg.TxHash.Hex(),
		lg.BlockNumber,
	); err != nil {
		return err
	}
	return nil
}

// AuctionCancelled event handle 拍卖取消事件处理
func (p *Processor) handleAuctionCancelled(lg *types.Log) error {
	//1. 检验topics
	if len(lg.Topics) < 2 {
		return fmt.Errorf("invalid AuctionCancelled topics len=%d", len(lg.Topics))
	}

	//2. 解析auctionId (唯一字段)
	auctionId := new(big.Int).SetBytes(lg.Topics[1].Bytes()).Uint64()

	//3. 调用repo 更新拍卖状态 "CANCELLED"
	return p.db.MarkAuctionCancelled(
		context.Background(),
		p.chainID,
		p.contract,
		auctionId,
		lg.TxHash.Hex(),
		lg.BlockNumber,
	)
}
