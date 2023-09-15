package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pusupalahemanthkumar/pkg/config"
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

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	fmt.Println("CreateBook")
	return b

}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	fmt.Println("Fetched All")
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	fmt.Println("Fetched By Id")
	return &getBook, db

}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("iD=?", Id).Delete(book)
	fmt.Println("Delete By Id")
	return book

}
