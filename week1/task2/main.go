package main

import "fmt"

// 指针题目1
func changeValue(a *int) int {
	*a = *a + 10
	// 返回指针 a 所指向的值，简称指针传递 引用传递函数（通过指针）
	return *a
}

// 指针题目2
func changeValue2(slice *[]int) *[]int {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] = (*slice)[i] * 2
	}

	return slice
}

func main() {
	//a := 1
	//取变量a的内存地址的值，值传递
	//fmt.Println(changeValue(&a))
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(changeValue2(&testSlice))
}
