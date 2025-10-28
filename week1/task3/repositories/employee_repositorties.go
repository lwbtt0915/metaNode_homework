package repositories

import (
	"dataDriver/database"
	"dataDriver/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository interface {
	GetByDepartment(department *models.Employees) ([]models.Employees, error)

	GetMaxSalary() (*models.Employees, error)
}

type employeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{
		db: database.GetDB(),
	}
}

// 取某个部门的所有员工信息，eg ""技术部" 的员工信息"
func (r *employeeRepository) GetByDepartment(emp *models.Employees) ([]models.Employees, error) {
	query := `SELECT id, name, department, salary FROM Employees where department =$1`

	var employees []models.Employees
	err := r.db.Select(&employees, query, emp.Department)
	if err != nil {
		return nil, fmt.Errorf("获取用户列表失败: %w", err)
	}

	return employees, nil
}

// employees 表中工资最高的员工信息，  todo 注意薪水怎样类型转换？
func (r *employeeRepository) GetMaxSalary() (*models.Employees, error) {
	query := `SELECT id, name, department, MAX(salary) FROM Employees `

	var employee models.Employees
	err := r.db.Get(&employee, query)

	if err != nil {
		return nil, fmt.Errorf("获取最高薪水的雇员失败: %w", err)
	}

	return &employee, nil
}
