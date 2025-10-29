package main

import (
	"dataDriver/database"
	"dataDriver/models"
	"dataDriver/repositories"
	"log"
)

//题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// todo 待测试
func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal("数据库初始化失败: %v", err)
	}

	defer database.CloseDB()

	employeeRepo := repositories.NewEmployeeRepository()

	employee, err := employeeRepo.GetMaxSalary()
	if err != nil {
		log.Printf("获取雇员最高薪水失败: %v", err)
	} else {
		log.Printf("用户: ID=%d, Name=%s, Sarlary=%s", employee.ID, employee.Name, employee.Salary)
	}

	param := &models.Employees{
		Department: "技术部",
	}

	ems, err := employeeRepo.GetByDepartment(param)

	if err != nil {
		log.Printf("获取雇员信息列表失败: %v", err)
	} else {
		log.Printf("雇员总数: %d", len(ems))
		for _, emp := range ems {
			log.Printf("用户: ID=%d, Name=%s, Department=%s,Salary=%s", emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	}

	bookRepos := repositories.NewBookRepository()
	books, err := bookRepos.GetBookByPrice()

	if err != nil {
		log.Printf("获取图书信息列表失败: %v", err)
	} else {
		log.Printf("图书总数: %d", len(books))
		for _, book := range books {
			log.Printf("图书: ID=%d, Author=%s, Price=%f,Title=%s", book.Id, book.Author, book.Price, book.Title)
		}
	}
}
