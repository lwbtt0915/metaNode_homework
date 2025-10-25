package main

import (
	"strconv"
)

// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}

	}
	return true
}

// //给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。,
func foundOnlyElment(arr []int) string {
	intMap := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		charstr := strconv.Itoa(arr[i])
		if _, exists := intMap[charstr]; !exists {
			intMap[charstr] = 1
		} else {
			intMap[charstr] = intMap[charstr] + 1
		}

	}

	for index, value := range intMap {
		if value == 1 {
			return index
		}
	}

	return ""
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

func isValid(s string) bool {
	stack := make([]byte, 0)
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		char := s[i]

		// 如果是右括号
		if matching, isRight := mapping[char]; isRight {
			// 检查栈是否为空
			if len(stack) == 0 {
				return false
			}
			// 获取栈顶元素
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 弹出栈顶

			// 检查是否匹配
			if top != matching {
				return false
			}
		} else {
			// 左括号，入栈
			stack = append(stack, char)
		}
	}

	// 最后栈应该为空
	return len(stack) == 0
}

// 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	for i := 0; i < len(digits); i++ {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
		}
	}

	return digits
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) (a int, b int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = i
				b = j
				return nums[i], nums[b]
			}
		}
	}

	return 0, 0
}

//查找字符串数组中的最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
func main() {
	// fmt.Println(IsPalindrome(111))
	// fmt.Println(IsPalindrome(-111))
	// fmt.Println(IsPalindrome(10))
	//var arr0 [5]int = [5]int{33, 33, 23, 23, 14}
	//fmt.Println(foundOnlyElment(arr0[:]))
	// fmt.Println(isValid("()[]{}"))
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 13))
	//fmt.Println(removeDuplicates([]int{2, 2, 3, 4, 7, 7, 7}))
	//fmt.Println(plusOne([]int{1, 2, 3, 4, 5}))

	//fmt.Println(longestCommonPrefix([]string{"abcde", "abce", "abrt"}))
}
