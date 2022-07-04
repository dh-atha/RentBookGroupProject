package main

import (
	"RentBookGroupProject/db"
	"fmt"
)

func main() {
	conn := db.InitDB()
	fmt.Println(conn.Ping())
}
