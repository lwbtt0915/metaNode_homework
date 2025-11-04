package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"web3/week1/task3/database"
	"web3/week1/task3/models"
)

type BookRepository interface {
	GetBookByPrice() ([]models.Book, error)
}

type bookRepository struct {
	db *sqlx.DB
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		db: database.GetDB(),
	}
}

func (r *bookRepository) GetBookByPrice() ([]models.Book, error) {
	//TODO implement me
	query := `SELECT id, title, author, price FROM book where price > $1`

	var book []models.Book
	err := r.db.Select(&book, query, 50)
	if err != nil {
		return nil, fmt.Errorf("获取用户列表失败: %w", err)
	}

	return book, nil
}
