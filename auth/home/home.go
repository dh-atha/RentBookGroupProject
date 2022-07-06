package home

import (
	"RentBookGroupProject/activities/book"
	"RentBookGroupProject/auth/login"
	"RentBookGroupProject/auth/register"
	"RentBookGroupProject/db"

	"fmt"
)

func HomeMenu() {
	conn := db.InitDB()
	db.MigrateDB(conn)

	var inputMenuAwal int
	fmt.Println("---Home Menu---")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. See Books")
	fmt.Print("\n")
	fmt.Print("99. Exit\n\n")
	fmt.Print("Input: ")
	fmt.Scanln(&inputMenuAwal)

	switch inputMenuAwal {
	case 1:
		register.Register(conn)
		HomeMenu()
	case 2:
		login.Login(conn)
	case 3:
		book.SeeBooks(conn)
		HomeMenu()
	case 99:
		fmt.Println("Exiting program...")
	default:
		fmt.Println("Menu tidak terdaftar")
	}
}
