// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./ITaxHandler.sol";

contract MyMeme is ERC20, Ownable {
    // 税收合约
    ITaxHandler public taxHandler;
    // 托管地址
    address public holder;
    // 单笔交易最大额度
    uint256 public maxTransfer;

    // 每日交易次数限制
    struct TradeCounter {
        uint256 day; // 当前日期
        uint8 count; // 今日交易次数
    }

    mapping(address => TradeCounter) tradeRecords;


    //  每日交易次数限制（最多3次）
    modifier checkDailyLimit() {
        uint256 today = block.timestamp / 1 days;
        TradeCounter storage record = tradeRecords[msg.sender];

        if (record.day < today) {
            record.day = today;
            record.count = 0;
        }

        require(record.count < 3, "exceeded daily tx limit");
        _;
        record.count += 1;
    }

    // 增加流动性，x * y = k, MME/ETH
    
    uint256 public totalLiquidity; //总流动性
    mapping(address => uint256) public liquidityBalance; // LP token，简易版 用户的LP代币余额
    uint256 public memeReserve;  //MME代币储备量
    uint256 public ethReserve;  // ETH 储备量

    constructor(uint256 _initSupply, address _holder) ERC20("MyMeme", "MME") Ownable(msg.sender) {
        holder = _holder;
        _mint(msg.sender, _initSupply);
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

   //设置最大交易额度
    function setMaxTransfer(uint256 _maxTransfer) public onlyOwner {
        maxTransfer = _maxTransfer;
    }

    function transfer(address to, uint256 value) public override checkDailyLimit returns (bool) {
        require(value <= maxTransfer, "exceed max transfer amount");
        address owner = _msgSender();

        // 增加交易税
        uint256 tax = taxHandler.getTax(value);
        uint256 taxedAmount = value - tax;

        // 转入税到托管账户
        if (tax > 0) {
            _transfer(owner, holder, tax);
        }
        _transfer(owner, to, taxedAmount);
        return true;
    }

    // 计算初始流动性
    // 求平方根
    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y > 3) {
            z = y;
            uint256 x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
    }

     //添加流动性
    function addLiquidity(uint256 memeAmount) external payable {
        require(memeAmount > 0 && msg.value > 0, "invalid amount");

        if (totalLiquidity == 0) {
            // 初始化流动性池子
            uint256 liquidity = sqrt(memeAmount * msg.value);
            require(liquidity > 0, "invalid initial liquidity");
            memeReserve = memeAmount;
            ethReserve = msg.value;
            totalLiquidity = liquidity;
            // 给流动性提供者发LP Token
            liquidityBalance[msg.sender] = liquidity;
        } else {
            uint256 mmeRatio = memeReserve * msg.value / ethReserve;
            require(memeAmount > mmeRatio, "MME amount not enough");

            uint256 liquidityMinted = (msg.value * totalLiquidity) / ethReserve;

            memeReserve += memeAmount;
            ethReserve += msg.value;
            totalLiquidity += liquidityMinted;
            liquidityBalance[msg.sender] += liquidityMinted;
        }

        _transfer(msg.sender, address(this), memeAmount);
    }


   // 移除流动性
    function removeLiquidity(uint256 liquidity) external {
        require(liquidity > 0, "invalid liquidity");
        require(liquidityBalance[msg.sender] > liquidity, "do not have enough liquidity");

        uint256 memeAmount = memeReserve * liquidity / totalLiquidity;
        uint256 ethAmount = ethReserve * liquidity / totalLiquidity;

        liquidityBalance[msg.sender] -= liquidity;
        totalLiquidity -= liquidity;
        memeReserve -= memeAmount;
        ethReserve -= ethReserve;

        _transfer(address(this), msg.sender, memeAmount);
        payable(msg.sender).transfer(ethAmount);
    }
}
