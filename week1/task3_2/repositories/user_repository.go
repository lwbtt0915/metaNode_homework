package repositories

import (
	"gorm.io/gorm"
)

type BlogService struct {
	db *gorm.DB
}
