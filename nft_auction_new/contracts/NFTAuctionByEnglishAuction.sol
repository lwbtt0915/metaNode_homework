// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// 导入ERC721接口
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "./interfaces/AggregatorV3Interface.sol";

/**
 * @title EnglishAuction
 * @dev 英式拍卖合约，支持出价递增拍卖
 * 特点：出价逐渐升高，价高者得
 */

 contract NFTAuctionByEnglishAuction is Initializable, OwnableUpgradeable, UUPSUpgradeable, ReentrancyGuardUpgradeable {
    
    //ETH 地址
    address public constant NATIVE_TOKEN= address(0);
    
     struct Auction {
         address  seller;  
        //address payable seller;      // 卖家地址（可支付）
        address nft;     // NFT 合约地址
        uint256 tokenId;         // NFT ID

        uint256 endTime;             // 拍卖结束时间（Unix时间戳）
      //  uint256 startTime;           // 拍卖开始时间（Unix时间戳）

        //最高出价原始金额
        uint256 highestBidAmount;          // 当前最高出价
        address highestBidToken;       // 当前最高出价者
       
        //最高出价(换算成USD)
        address highestBidder;
        uint256 highestBidUsd18;
        
        uint256 minBidUsd18;   // 起拍价（wei）

        bool ended; //拍卖是否结束
     }

     uint256 public nextAuctionId; //新创建的拍卖分配唯一的拍卖 ID
     mapping(uint256 => Auction) public auctions; // mapping 容器存拍卖 ID（键）关联到对应的拍卖详情
    

    // 只支持一种 ERC20：USDC
    IERC20 public usdc;

    //Chainlink feeds
    AggregatorV3Interface public ethUsdFeed; //获取 ETH 对 USD 的实时价格，用于「ETH 与 USDC 之间的计价换算」
    AggregatorV3Interface public usdcUsdFeed; //获取 USDC 对 USD 的实时价格
    

    address public feeRecipient;//合约收取的所有手续费（USDC/ETH）都会转账到这个地址
    uint256 public feeBps; //定义合约收取的手续费比例


    //卖家创建拍卖
    event CreateAuction(
        uint256 indexed auctionId,
        address indexed seller,
        address indexed nft,
        uint256 tokenId,
        uint256 endTime,
        uint256 minBidUsd18
    );

    // 买家出价
    event BidAuction(
        uint256 indexed auctionId,
        address indexed bidder,
        address bidToken,
        uint256 bidAmount,
        uint256 bidUsd18
    );

    //拍卖结束
    event AuctionEnded(
        uint256 indexed auctionId,
        address winner,
        address payToken,
        uint256 payAmount,
        uint256 payUsd18
    );


    // 取消拍卖
    event AuctionCancelled(uint256 indexed auctionId);


     // @custom:oz-upgrades-unsafe-allow constructor
    // constructor() {
    //     //// 关键：防止初始化函数被调用
    //     _disableInitializers();
    // }

    //升级合约
    function initialize(
        address _ethUsdFeed,
        address _usdc,
        address _usdcUsdFeed,
        address _feeRecipient,
        uint256 _feeBps
    ) external initializer {
      __Ownable_init(msg.sender); //初始化所有者
      __UUPSUpgradeable_init();  // 初始化UUPS升级逻辑
      __ReentrancyGuard_init(); // 初始化防重入保护

      ethUsdFeed=AggregatorV3Interface(_ethUsdFeed);
      usdc = IERC20(_usdc);
      usdcUsdFeed = AggregatorV3Interface(_usdcUsdFeed);

      feeRecipient = _feeRecipient;
      feeBps= _feeBps;
    }

   //逻辑合约的 _authorizeUpgrade 函数
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

   //创建拍卖
    function createAuction(
        address nft,
        uint256 tokenId,
        uint256 durationSeconds,
        uint256 minBidUsd18
    ) external nonReentrant returns (uint256 auctionId) {
        require(durationSeconds > 0, "Bad duration");
        require(minBidUsd18 > 0, "Bad minBidUsd");

        auctionId = ++nextAuctionId;
        auctions[auctionId] = Auction({
             seller: msg.sender,
             nft: nft,
             tokenId: tokenId,
             endTime: block.timestamp + durationSeconds,
             ended: false,
             highestBidToken: NATIVE_TOKEN,
             highestBidAmount: 0,
             highestBidUsd18: 0,
             highestBidder: address(0),
             minBidUsd18:minBidUsd18
        });


        //托管NFT
    IERC721(nft).transferFrom(msg.sender, address(this), tokenId);

    //触发创建拍卖事件
    emit CreateAuction(auctionId, msg.sender, nft, tokenId, auctions[auctionId].endTime, minBidUsd18);
    }

   /* ---------------- Oracle helpers (USD 18 decimals) ---------------- */
   //将任意小数位的数值统一转换为 18 位小数
   function _scaleTo18(uint256 value, uint8 decimals_) internal pure returns (uint256) {
        if (decimals_ == 18) return value;
        if (decimals_ < 18) return value * (10 ** (18 - decimals_));
        return value / (10 ** (decimals_ - 18));
   }



   //读取 Chainlink 价格喂价
   function _readFeed(AggregatorV3Interface feed) internal view returns (uint256 answer, uint8 dec) {
        (, int256 a,, uint256 updatedAt,) = feed.latestRoundData();
        require(a > 0, "Oracle: bad answer");
        require(updatedAt > 0, "Oracle: stale");
        answer = uint256(a);
        dec = feed.decimals();
    }



    //ETH 金额 → 18 位 USD 金额
    function _ethToUsd18(uint256 weiAmount) internal view returns (uint256) {
        // ===== 本地链 / Hardhat：禁用 Oracle =====
        if (block.chainid == 31337) {
            // mock：1 ETH = 2000 USD（18位）
            // wei(1e18) * 2000 = 2000e18
            return weiAmount * 2000;
        }

        // ===== 测试网 / 主网：真实 Oracle =====
        (uint256 price, uint8 dec) = _readFeed(ethUsdFeed);
        uint256 usd = (weiAmount * price) / 1e18;
        return _scaleTo18(usd, dec);
    }



    //USDC 金额 → 18 位 USD 金额
    function _usdcToUsd18(uint256 usdcAmount) internal view returns (uint256) {
        if (block.chainid == 31337) {
            // mock：1 USDC = 1 USD
            // usdc 通常 6 位
            return usdcAmount * 1e12;
        }

        (uint256 price, uint8 feedDec) = _readFeed(usdcUsdFeed);
        uint256 usdWithFeedDec = (usdcAmount * price) / 1e6;
        return _scaleTo18(usdWithFeedDec, feedDec);
    }



    //ETH 出价逻辑
    function bidEth(uint256 auctionId) external payable nonReentrant {
        Auction storage a = auctions[auctionId];

        require(a.seller !=address(0), "No Auction");
        require(!a.ended, "Ended");
        require(block.timestamp < a.endTime, "Auction over");
        require(msg.value > 0, "Zero bid");

        uint256 bidUsd18 = _ethToUsd18(msg.value);
        _placeBid(a, auctionId, msg.sender, NATIVE_TOKEN, msg.value, bidUsd18);
    }
    

   //USDC 出价逻辑
    function bidUSDC(uint256 auctionId, uint256 usdcAmount) external nonReentrant {
        Auction storage a = auctions[auctionId];
        require(a.seller != address(0), "No auction");
        require(!a.ended, "Ended");
        require(block.timestamp < a.endTime, "Auction over");
        require(usdcAmount > 0, "Zero bid");

        uint256 bidUsd18 = _usdcToUsd18(usdcAmount);

        require(usdc.transferFrom(msg.sender, address(this), usdcAmount), "USDC transferFrom failed");
        _placeBid(a, auctionId, msg.sender, address(usdc), usdcAmount, bidUsd18);
    }


    function _placeBid(
        Auction storage a,
        uint256 auctionId,
        address bidder,
        address bidToken,
        uint256 bidAmount,
        uint256 bidUsd18
    ) internal {
        require(bidUsd18 >= a.minBidUsd18, "Below minBidUsd");
        require(bidUsd18 > a.highestBidUsd18, "Bid not high enough");

        // 退还上一位最高出价者（按原币种原金额退）
        if (a.highestBidder != address(0)) {
            _refund(a.highestBidder, a.highestBidToken, a.highestBidAmount);
        }

        a.highestBidder = bidder;
        a.highestBidToken = bidToken;
        a.highestBidAmount = bidAmount;
        a.highestBidUsd18 = bidUsd18;

        emit BidAuction(auctionId, bidder, bidToken, bidAmount, bidUsd18);
    }


   //退还上个出价者
    function _refund(address to, address token, uint256 amount) internal {
        if (amount == 0) return;
        if (token == NATIVE_TOKEN) {
            (bool ok,) = payable(to).call{value: amount}("");
            require(ok, "Refund ETH failed");
        } else {
            require(IERC20(token).transfer(to, amount), "Refund ERC20 failed");
        }
    }


    //结束拍卖
    function endAuction(uint256 auctionId) external nonReentrant {
        Auction storage a = auctions[auctionId];
        require(a.seller != address(0), "No auction");
        require(!a.ended, "Already ended");
        require(block.timestamp >= a.endTime, "Not ended yet");

        a.ended = true;

        if (a.highestBidder == address(0)) {
            // 无人出价：退回 NFT
            IERC721(a.nft).transferFrom(address(this), a.seller, a.tokenId);
            emit AuctionCancelled(auctionId);
            return;
        }

        // NFT 给赢家
        IERC721(a.nft).transferFrom(address(this), a.highestBidder, a.tokenId);

        // 资金给卖家（赢家用什么付，就给什么）
        if (a.highestBidToken == NATIVE_TOKEN) {
            (bool ok,) = payable(a.seller).call{value: a.highestBidAmount}("");
            require(ok, "Pay seller ETH failed");
        } else {
            require(IERC20(a.highestBidToken).transfer(a.seller, a.highestBidAmount), "Pay seller ERC20 failed");
        }

        emit AuctionEnded(auctionId, a.highestBidder, a.highestBidToken, a.highestBidAmount, a.highestBidUsd18);
    }


    //获取拍卖信息
    function getAuction(uint256 auctionId) external view returns (Auction memory) {
        return auctions[auctionId];
    }

    receive() external payable {}
 }

