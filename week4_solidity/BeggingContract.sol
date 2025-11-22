// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract BeggingContract {
    //捐赠者对应的捐赠金额
    mapping(address pepole => uint256 value)  donations;
    //合约所有者地址
    address public owner;
    //总捐赠金额
    uint256 public totalDonations;


    constructor(){
       owner = msg.sender;
    }


    event Donated(address indexed donor, uint256 amount);
    event Withdrawn(address indexed owner, uint256 amount);

    // 修饰符：只有所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }


    // donate 函数：允许用户捐赠以太币
    function donate() external payable {
        require(msg.value > 0, "Donation amount must be greater than 0");
        donations[msg.sender] +=msg.value;
        totalDonations+=msg.value;

        emit Donated(msg.sender, msg.value);
    }


    //允许所有者提取所有资金
    function withdraw() external onlyOwner{
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        payable(owner).transfer(balance);
        emit Withdrawn(owner, balance);
    }

    //查询某个地址的捐款金额
    function getContract(address donor) external view returns(uint256) {
        return donations[donor];
    }

    //获取合约当前金额
    function getContractBalance() external view returns(uint256) {
        return address(this).balance;
    }


    //获取所有捐赠人
    function getTotalDonors() external view returns(uint256) {
        return totalDonations;
    }
}