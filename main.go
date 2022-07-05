package main

import (
	"RentBookGroupProject/db"
	"fmt"
)

func main() {
	conn := db.InitDB()
	db.MigrateDB(conn)

	var inputMenuAwal int

	for inputMenuAwal != 99 {
		fmt.Print("\n")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. See Books")
		fmt.Print("\n")
		fmt.Print("99. Exit\n\n")
		fmt.Print("Input: ")
		fmt.Scanln(&inputMenuAwal)

		switch inputMenuAwal {
		case 1:
			Register()
		case 2:
			Login()
		case 3:
			SeeBooks()
		case 99:
			fmt.Println("Exiting program...")
		default:
			fmt.Println("Menu tidak terdaftar")
		}
	}
}

func Register() {}

func Login() {}

func SeeBooks() {}
