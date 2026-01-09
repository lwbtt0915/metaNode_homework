// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./NFTAuctionByEnglishAuction.sol";

contract NFTAuctionByEnglishAuctionV2 is NFTAuctionByEnglishAuction {
    function version() external pure returns (uint256) {
        return 2;
    }

    uint256 public minBidAmount; // V2 新增


    // ===== 2. 核心修复：添加 initializer 函数 =====
    // 即使 V2 不修改初始化逻辑，也必须保留该函数（兼容代理初始化）
    function initialize(
    ) external initializer {
        // 第一步：调用所有父合约的初始化函数（必须！）
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();
        __UUPSUpgradeable_init();
    }
}