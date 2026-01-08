// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface AggregatorV3Interface {
    //读取喂价合约元数据
    function decimals()external view returns(uint8);
    //回价格喂价的描述信息（字符串），用于标识喂价类型，比如 ETH/USD 喂价合约返回
    function description()external view returns(string memory);
    //返回价格聚合器的版本号（如 3），用于兼容不同版本的喂价接口
    function version()external view returns(uint256);


    //历史价格数据函数
    function getRoundData(uint80 _roundId) external view returns(
    uint80  roundId,int256 answer, uint256 startAt,uint256 updateAt, uint80 answeredInRound);

   //最新价格数据函数
    function latestRoundData() external view returns (
        uint80 roundId,int256 answer, uint256 startAt, uint256 updateAt, uint80 answeredInRound);

}