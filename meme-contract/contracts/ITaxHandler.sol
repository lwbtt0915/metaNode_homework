// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//税率接口
interface ITaxHandler {
    function getTax(uint256 amount) external returns (uint256);
}
