package dashboard

import (
	"RentBookGroupProject/activities/book"
	"RentBookGroupProject/activities/rent"
	"RentBookGroupProject/activities/user"
	"RentBookGroupProject/db"

	"fmt"
)

func DashboardMenu() {
	conn := db.InitDB()
	db.MigrateDB(conn)

	var inputMenu int
	fmt.Println("\n---Dashboard---")
	fmt.Println("1. Profile")
	fmt.Println("2. My Own Books")
	fmt.Println("3. Book I Rented")
	fmt.Println("4. Return Books")
	fmt.Print("\n")
	fmt.Print("99. Exit\n\n")
	fmt.Print("Input: ")
	fmt.Scanln(&inputMenu)

	switch inputMenu {
	case 1:
		user.SeeProfile(conn)
	case 2:
		book.MyBooks()
	case 3:
		rent.BooksRented()
	case 4:
		rent.ReturnBook()

	case 99:
		fmt.Println("Exiting program...")
	default:
		fmt.Println("Menu tidak terdaftar")
		DashboardMenu()
	}
}
