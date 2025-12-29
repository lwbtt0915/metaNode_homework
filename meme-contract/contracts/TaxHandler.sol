// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ITaxHandler.sol";

contract TaxHandler is ITaxHandler {
    //阶梯税率规则
    struct stepTax {
        uint256 level; //阶梯的金额门槛
        uint8 rate; // 对应阶梯的税率
    }

    //动态数组，存储所有阶梯税率规则
    stepTax[] public stepTaxes;

    // 设置阶梯税率
    function setStepTaxes(uint256[] memory levels, uint8[] memory rates) external {
        require(levels.length > 0);
        require(levels.length == rates.length, "length must same");
        for (uint256 i = 0; i < levels.length; i++) {
            stepTax memory t = stepTax({level: levels[i], rate: rates[i]});
            stepTaxes.push(t);
        }
    }

     
     // 计算税费
    function getTax(uint256 amount) external view returns (uint256) {
        for (uint256 i = stepTaxes.length; i > 0; i--) {
            if (amount >= stepTaxes[i - 1].level) {
                return amount * stepTaxes[i - 1].rate / 100;
            }
        }
        return 0;
    }
}
