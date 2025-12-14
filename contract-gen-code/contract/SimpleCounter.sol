
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleCounter {
    uint256 private count;
    
    // 获取当前计数值
    function getCount() public view returns (uint256) {
        return count;
    }
    
    // 增加计数器
    function increment() public {
        count += 1;
    }
    
    // 减少计数器
    function decrement() public {
        require(count > 0, "Counter: cannot decrement below zero");
        count -= 1;
    }
    
    // 重置计数器
    function reset() public {
        count = 0;
    }
}
















