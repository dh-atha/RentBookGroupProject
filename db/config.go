package db

import (
	"RentBookGroupProject/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/rentbookgroupproject?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entities.User{})
	conn.AutoMigrate(entities.Book{})
	conn.AutoMigrate(entities.Rent{})
}
