package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/RentBookGroupProject")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
