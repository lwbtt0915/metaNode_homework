// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./NFTAuctionByEnglishAuction.sol";

contract NFTAuctionByEnglishAuctionV2 is NFTAuctionByEnglishAuction {
    function version() external pure returns (uint256) {
        return 2;
    }
}