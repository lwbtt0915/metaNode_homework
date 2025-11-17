// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract MergeSortedArrays {


    /**
     * @dev 合并两个升序有序数组为一个升序有序数组
     * @param arr1 第一个升序有序数组（支持空数组）
     * @param arr2 第二个升序有序数组（支持空数组）
     * @return 合并后的升序有序数组
     * @notice 输入数组需保证升序有序，否则结果不保证正确性
     */
    function merge( uint256[] calldata arr1,uint256[] calldata arr2) external pure returns (uint256[] memory) {
        uint256 n = arr1.length;
        uint256 m = arr2.length;
        uint256[] memory result = new uint256[](n + m); // 预分配结果数组空间（gas优化）

        uint256 i = 0; // arr1 指针
        uint256 j = 0; // arr2 指针
        uint256 k = 0; // result 指针

        // 双指针遍历，比较并合并
        while (i < n && j < m) {
            if (arr1[i] <= arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 追加 arr1 剩余元素
        while (i < n) {
            result[k] = arr1[i];
            i++;
            k++;
        }

        // 追加 arr2 剩余元素
        while (j < m) {
            result[k] = arr2[j];
            j++;
            k++;
        }

        return result;
    }

    /**
     * @dev 重载：支持合并两个 int256 类型的升序有序数组
     * @notice 适配有符号整数场景
     */
    function merge(
        int256[] calldata arr1,
        int256[] calldata arr2
    ) external pure returns (int256[] memory) {
        uint256 n = arr1.length;
        uint256 m = arr2.length;
        int256[] memory result = new int256[](n + m);

        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        while (i < n && j < m) {
            if (arr1[i] <= arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }

        while (i < n) {
            result[k] = arr1[i];
            i++;
            k++;
        }

        while (j < m) {
            result[k] = arr2[j];
            j++;
            k++;
        }

        return result;
    }
}
