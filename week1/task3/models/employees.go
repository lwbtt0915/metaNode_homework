package models

type Employees struct {
	ID         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Department string `db:"department" json:"department"`
	Salary     string `db:"salary" json:"salary"`
}
