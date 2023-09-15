package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "springstudent:springstudent@tcp(127.0.0.1:3306)/goBookStore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connection Done !!")
	db = d

}

func GetDB() *gorm.DB {
	return db
}
