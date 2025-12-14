// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20; // 推荐使用较新的编译器版本

contract SimpleCounter {
    // 状态变量
    uint256 private count;
    address public owner; // 合约拥有者地址

    // 事件：记录计数器变化（便于前端监听）
    event CountIncremented(uint256 newCount);
    event CountDecremented(uint256 newCount);
    event CountReset(uint256 newCount);
    event CountIncreasedByStep(uint256 step, uint256 newCount);
    event CountDecreasedByStep(uint256 step, uint256 newCount);

    // 修饰器：限制仅合约拥有者可调用（解决权限问题）
    modifier onlyOwner() {
        require(msg.sender == owner, "SimpleCounter: caller is not the owner");
        _;
    }

    // 构造函数：初始化拥有者为部署者
    constructor() {
        owner = msg.sender;
        count = 0; // 显式初始化（可选，默认值为0）
    }

    // 获取当前计数值
    function getCount() public view returns (uint256) {
        return count;
    }

    // 增加计数器（+1）
    function increment() public onlyOwner {
        count += 1;
        emit CountIncremented(count); // 触发事件
    }

    // 减少计数器（-1）：优化异常提示，仅Owner可调用
    function decrement() public onlyOwner {
        if (count == 0) {
            revert("SimpleCounter: count is already zero (cannot decrement)");
        }
        count -= 1;
        emit CountDecremented(count);
    }

    // 重置计数器：仅Owner可调用
    function reset() public onlyOwner {
        count = 0;
        emit CountReset(count);
    }

    // 扩展功能：自定义步长增加（防止溢出）
    function increaseByStep(uint256 step) public onlyOwner {
        require(step > 0, "SimpleCounter: step must be greater than zero");
        // 使用unchecked仅在确认无溢出时使用，或用safeAdd（推荐）
        count = count + step; // Solidity 0.8+默认溢出检查，若溢出会revert
        emit CountIncreasedByStep(step, count);
    }

    // 扩展功能：自定义步长减少（防止下溢）
    function decreaseByStep(uint256 step) public onlyOwner {
        require(step > 0, "SimpleCounter: step must be greater than zero");
        require(count >= step, "SimpleCounter: step exceeds current count");
        count -= step;
        emit CountDecreasedByStep(step, count);
    }

    // 可选：转移合约拥有权
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0), "SimpleCounter: new owner is the zero address");
        owner = newOwner;
    }
}