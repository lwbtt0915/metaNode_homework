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

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person     Person
}

func (m Employee) PrintInfo() {
	id := m.EmployeeID
	fmt.Println("id", id)
	name := m.Person.Name
	fmt.Println("name", name)
	age := m.Person.Age
	fmt.Println("age", age)
}

//todo ✅Channel
//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信。

func Communication() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}

		fmt.Println("发送完成")
	}()

	go func() {
		for value := range ch {
			fmt.Printf("接收: %d\n", value)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("接收完成")
	}()

	time.Sleep(time.Second * 6)
	fmt.Println("程序结束")
}

//todo 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

func BufferChannelDemo() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}

		close(ch)
	}()

	go func() {
		for value := range ch {
			fmt.Printf("接收: %d\n", value)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second * 20)
	fmt.Println("程序结束")
}

//todo ✅锁机制
//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ： sync.Mutex 的使用、并发数据安全。

//todo 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全。

func main() {
	//a := 1
	//取变量a的内存地址的值，值传递
	//fmt.Println(changeValue(&a))
	//testSlice := []int{1, 2, 3, 4, 5}
	//fmt.Println(changeValue2(&testSlice))
	//printNum()

	//r := Rectangle{1, 2}
	//r.Area()
	//r.Perimeter()
	//
	//c := Circle{1}
	//c.Area()
	//c.Perimeter()

	//e := Employee{
	//	EmployeeID: 1,
	//	Person:     Person{"lwb", 3},
	//}
	//
	//e.PrintInfo()

	//Communication()

	BufferChannelDemo()
}
