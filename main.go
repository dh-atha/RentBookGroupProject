package main

import (
	"RentBookGroupProject/db"
	"RentBookGroupProject/entity"
	"fmt"
	"log"
)

func main() {
	conn := db.InitDB()
	// fmt.Println(conn.Ping())

	var usersData entity.Users
	var daftarUsers = []entity.Users{}

	rows, err := conn.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&usersData)
		if err != nil {
			log.Fatal(err)
		}
		daftarUsers = append(daftarUsers, usersData)
	}
	fmt.Println(daftarUsers)
}
