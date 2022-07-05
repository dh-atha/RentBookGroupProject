package main

import (
	"RentBookGroupProject/db"
	"log"
)

func main() {
	conn := db.InitDB()
	log.Println(conn)
}
