package indexer

import (
	"backend/internal/config"
	"backend/internal/db"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"time"
)

const CursorName = "auction_market_cursor"

// 索引器结构体
type Indexer struct {
	cfg       *config.Config
	db        *db.MySQL
	client    *ethclient.Client
	contract  common.Address
	Processor *Processor
}

// 创建索引器
func NewIndexer(cfg *config.Config, db *db.MySQL) (*Indexer, error) {
	//连接以太坊
	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(cfg.AuctionContract)

	return &Indexer{
		cfg:       cfg,
		db:        db,
		client:    client,
		contract:  contractAddr,
		Processor: NewProcessor(db, cfg.ChainID, strings.ToLower(contractAddr.Hex())),
	}, nil
}

// 启动索引器
func (i *Indexer) Start(ctx context.Context) {
	go i.loop(ctx)
}

// 持续轮询指定区块范围
func (i *Indexer) loop(ctx context.Context) {
	log.Printf(
		"[indexer] started chain_id=%d rpc=%s contract=%s",
		i.cfg.ChainID,
		i.cfg.RpcURL,
		i.contract.Hex(),
	)

	ticker := time.NewTicker(i.cfg.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[indexer] stopped")
			return

		case <-ticker.C:
			if err := i.tick(ctx); err != nil {
				log.Printf("[indexer] tick error: %v", err)
			}
		}

	}
}

func (i *Indexer) tick(ctx context.Context) error {
	chainID := i.cfg.ChainID
	contract := strings.ToLower(i.contract.Hex())

	// 从数据库获取游标：查询「该链ID+该合约+该游标名」对应的最后扫描区块高度
	last, err := i.db.GetCursor(chainID, contract, CursorName)

	if err != nil {
		log.Printf("[indexer][debug] GetCursor error: %v", err)

		// 如果是「游标不存在」（首次启动）：初始化游标
		if err != db.ErrCursorNotFound {
			return err
		}

		// 获取当前链上最新区块号：ethclient.BlockNumber 是 go-ethereum 中查询最新区块高度的核心方法
		latest, err := i.client.BlockNumber(ctx)
		if err != nil {
			return err
		}

		// 初始化扫描起始高度：如果最新区块号 > 初始扫描深度（InitScanDepth），就从 最新 - InitScanDepth 开始
		// 比如 InitScanDepth=1000，最新区块是 10000，就从 9000 开始扫描（避免首次启动扫描全量区块）
		start := latest
		log.Print("[indexer][debug] latest block number: ", latest)
		if latest > i.cfg.InitScanDepth {
			start = latest - i.cfg.InitScanDepth
		}

		// 把初始游标存入数据库：下次扫描从 start 开始
		return i.db.UpdateCursor(chainID, contract, CursorName, start)
	}

	// 最新的区块
	latest, err := i.client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	// 处理区块确认数：如果配置了 Confirmations（比如6），就取 latest - 6 作为扫描上限
	// 原因：以太坊区块可能会有重组，等待 N 个确认后再扫描，保证数据最终性
	if i.cfg.Confirmations > 0 && latest > i.cfg.Confirmations {
		latest -= i.cfg.Confirmations
	}

	// 如果游标已经追上最新区块：无需扫描，直接返回
	if last >= latest {
		return nil
	}

	//3.控制batch（防Rpc 限制）
	//本次扫描的起始区块：游标+1（增量扫描）
	from := last + 1

	// 本次扫描的结束区块：起始+批量大小（BatchSize），避免单次扫描太多区块触发 RPC 节点限流
	to := from + i.cfg.BatchSize
	if to > latest {
		to = latest
	}

	log.Printf("[indexer] scan blocks [%d → %d] (latest=%d)", from, to, latest)

	//4. 构造filter
	ab := mustParseABI()
	topic := []common.Hash{
		ab.Events["CreateAuction"].ID,
		ab.Events["BidAuction"].ID,
		ab.Events["AuctionEnded"].ID,
		ab.Events["AuctionCancelled"].ID}

	filter := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{i.contract},
		Topics:    [][]common.Hash{topic},
	}

	logs, err := i.client.FilterLogs(ctx, filter)

	if err != nil {
		return err
	}

	// 5. 处理log日志
	for _, log1 := range logs {
		if err := i.Processor.handleLog(ctx, log1); err != nil {
			log.Printf(
				"[indexer] handle log error block=%d tx=%s err=%v",
				log1.BlockNumber,  //所在区块号
				log1.TxHash.Hex(), //交易哈希
				err)
		}
	}

	//6. 处理游标
	if err := i.db.UpdateCursor(chainID, contract, CursorName, to); err != nil {
		return err
	}
	log.Printf("[indexer] updated cursor to %d", to)

	return nil
}
