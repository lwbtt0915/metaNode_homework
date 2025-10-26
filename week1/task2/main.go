package main

import (
	"fmt"
	"time"
)

// 指针题目1
func changeValue(a *int) int {
	*a = *a + 10
	// 返回指针 a 所指向的值，简称指针传递 引用传递函数（通过指针）
	return *a
}

// 指针题目2 实现一个函数 接收一个整数切片的指针，将切片中的每个元素乘以2
func changeValue2(slice *[]int) *[]int {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] = (*slice)[i] * 2
	}

	return slice
}

// goroutine 题目一
func printNum() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println("goroutine1 奇数 ", i)
			}
		}
	}()

	go func() {
		for j := 0; j < 10; j++ {
			if j%2 == 0 {
				fmt.Println("goroutine2 偶数 ", j)
			}
		}
	}()

	time.Sleep(time.Millisecond * 100)
}

//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

// Task 任务结构体
type Task struct {
	Name string
	Func func() error
}

// 任务结束结果 结构体
type TaskResult struct {
	TaskName string
	Duration time.Duration
	Err      error
}

// 调度器
func TaskScheduler(tasks []Task) []TaskResult {
	results := make(chan TaskResult, len(tasks))

	for _, task := range tasks {
		go func(t Task) {
			startTime := time.Now()
			err := t.Func()
			duration := time.Since(startTime)

			results <- TaskResult{
				TaskName: t.Name,
				Duration: duration,
				Err:      err,
			}
		}(task)
	}

	// 收集所有任务结果
	var taskResults []TaskResult
	for i := 0; i < len(tasks); i++ {
		result := <-results
		taskResults = append(taskResults, result)
		fmt.Printf("任务 '%s' 执行完成，耗时: %v，错误: %v\n", result.TaskName, result.Duration, result.Err)
	}

	return taskResults
}

//面向对象
//题目一：   定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
//然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area()      //面积
	Perimeter() // 周长
}

// 长方形
type Rectangle struct {
	a int
	b int
}

func (r Rectangle) Area() {
	a := r.a * r.b
	fmt.Println("Rectangle Area", a)
}

func (r Rectangle) Perimeter() {
	a := r.a*2 + r.b*2
	fmt.Println("Rectangle Perimeter", a)
}

// 圆形
type Circle struct {
	r float64
}

func (c Circle) Area() {
	a := 3.14 * c.r * c.r
	fmt.Println("Circle Area", a)
}

func (c Circle) Perimeter() {
	a := 2 * 3.14 * c.r
	fmt.Println("Circle Perimeter", a)
}

func main() {
	//a := 1
	//取变量a的内存地址的值，值传递
	//fmt.Println(changeValue(&a))
	//testSlice := []int{1, 2, 3, 4, 5}
	//fmt.Println(changeValue2(&testSlice))
	//printNum()

	r := Rectangle{1, 2}
	r.Area()
	r.Perimeter()

	c := Circle{1}
	c.Area()
	c.Perimeter()
}
