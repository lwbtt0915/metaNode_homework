// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract IntegerToRoman {

uint256[] private values = [1000,900,500,400,100,90,50,40,10,9,5,4,1];

string[] private symbols = ["M", "CM", "D", "CD",
"C", "XC", "L", "XL",
"X", "IX", "V", "IV",
"I"];


function intToRoman(uint256 num) public view returns (string memory) {
// 校验输入范围（罗马数字不支持 0 和大于 3999 的数）
require(num >= 1 && num <= 3999, "Number must be between 1 and 3999");
// 字符串拼接使用 bytes 更高效（Solidity 中 string 本质是 bytes）
bytes memory result = new bytes(0);

// 从大到小遍历映射表
for (uint256 i = 0; i < values.length; i++) {
// 当当前数值 <= 剩余整数时，拼接对应罗马字符并减去数值
while (num >= values[i]) {
// 将罗马字符（string）转换为 bytes 拼接到结果中
result = abi.encodePacked(result, symbols[i]);
num -= values[i];
}

// 提前退出循环（优化：当 num 减为 0 时无需继续遍历）
if (num == 0) break;
}

// 将 bytes 转换为 string 返回
return string(result);
}
}