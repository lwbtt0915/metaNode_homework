// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract RomanToInteger {

    uint256[] private values = [1000,900,500,400,100,90,50,40,10,9,5,4,1];

    string[] private symbols = ["M", "CM", "D", "CD","C", "XC", "L", "XL","X", "IX", "V", "IV", "I"];

    function romanToInt(string calldata roman) external pure returns (uint256) {
        uint256 length = bytes(roman).length;
        require(length > 0, "Empty string is invalid");

        uint256 result = 0;
        uint256 previousValue = 0;

        // 从右往左遍历（更高效的处理方式）
        for (uint256 i = length - 1; i < length; i--) {
            uint256 currentValue = getValue(bytes(roman)[i]);
            // 规则：如果当前值小于前一个值，说明是减法组合（如 IV=4, IX=9 等）
            if (currentValue < previousValue) {
                result -= currentValue;
            } else {
                result += currentValue;
            }

            previousValue = currentValue;
        }

        // 验证结果是否在有效范围内
        require(result >= 1 && result <= 3999, "Result out of valid range (1-3999)");
        return result;
    }


    function getValue(bytes1 c) private pure returns (uint256) {
        if (c == 'I') return 1;
        if (c == 'V') return 5;
        if (c == 'X') return 10;
        if (c == 'L') return 50;
        if (c == 'C') return 100;
        if (c == 'D') return 500;
        if (c == 'M') return 1000;
        revert("Invalid Roman character");
    }

}