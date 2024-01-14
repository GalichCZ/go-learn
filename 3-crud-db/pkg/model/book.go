package model

import (
	"crud-db/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	tx := db.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(b).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	defer tx.Commit()

	return b, nil
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) (Book, error) {
	tx := db.Begin()

	if tx.Error != nil {
		return Book{}, tx.Error
	}

	var deletedBook Book

	if err := tx.Where("ID = ?", Id).First(&deletedBook).Error; err != nil {
		tx.Rollback()
		return Book{}, err
	}

	if err := tx.Delete(&deletedBook).Error; err != nil {
		tx.Rollback()
		return Book{}, err
	}

	tx.Commit()

	return deletedBook, nil
}
