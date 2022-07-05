package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/rentbookgroupproject"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
