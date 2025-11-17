// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {


    /**
     * @dev 二分查找：在升序有序 uint256 数组中精确查找目标值
     * @param sortedArr 升序有序数组（支持空数组、重复元素）
     * @param target 待查找的目标值（非负整数）
     * @return 找到则返回目标值的索引；未找到返回 type(uint256).max（无效标识）
     * @notice 输入数组必须保证升序有序，否则结果不保证正确
     */
    function binarySearchUint(uint256[] calldata sortedArr,uint256 target) external pure returns (uint256) {
        uint256 arrLen = sortedArr.length;
        if (arrLen == 0) return type(uint256).max;

        uint256 left = 0;
        uint256 right = arrLen - 1;

        while (left <= right) {
            // 计算 mid：left + (right - left)/2 避免 (left+right) 溢出
            uint256 mid = left + (right - left) / 2;
            uint256 midValue = sortedArr[mid];

            if (midValue == target) {
                return mid;
            } else if (midValue < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return type(uint256).max;
    }

}