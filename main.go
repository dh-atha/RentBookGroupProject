package main

import (
	"RentBookGroupProject/db"
	"RentBookGroupProject/entities/book"
	"RentBookGroupProject/entities/user"
	"fmt"
)

func main() {
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

	for inputMenuAwal != 99 {
		switch inputMenuAwal {
		case 1:
			user.Register(conn)
		case 2:
			user.Login(conn)
			// var inputMenuDashboard int
			// fmt.Println("\n---Dashboard---")
			// fmt.Println("1. Profile")
			// fmt.Println("2. My Own Books")
			// fmt.Println("3. Book I Rented")
			// fmt.Println("4. Return Books")
			// fmt.Print("\n")
			// fmt.Print("99. Exit\n\n")
			// fmt.Print("Input: ")
			// fmt.Scanln(&inputMenuDashboard)

			// for inputMenuDashboard != 99 {
			// 	switch inputMenuDashboard {
			// 	case 1:
			// 		user.SeeProfile(conn)
			// 	case 2:
			// 		book.MyBooks()
			// 	case 3:
			// 		rent.BooksRented()
			// 	case 4:
			// 		rent.ReturnBook()
			// 	case 99:
			// 		break
			// 	default:
			// 		fmt.Println("Menu tidak terdaftar")
			// 	}
			// }
		case 3:
			book.SeeBooks(conn)
		case 99:
			fmt.Println("Exiting program...")
		default:
			fmt.Println("Menu tidak terdaftar")
		}
	}
}
