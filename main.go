package main

import (
	"RentBookGroupProject/db"
	"RentBookGroupProject/entities"
	"fmt"
)

func main() {
	conn := db.InitDB()
	db.MigrateDB(conn)

	var inputMenuAwal int

	for inputMenuAwal != 99 {
		fmt.Println("\n\t---Home Menu---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. See Books")
		fmt.Print("\n")
		fmt.Print("99. Exit\n\n")
		fmt.Print("Input: ")
		fmt.Scanln(&inputMenuAwal)

		switch inputMenuAwal {
		case 1:
			entities.Register(conn)
		case 2:
			check := entities.Login(conn)
			if check {
				var inputMenuDashboard int
				for inputMenuDashboard != 99 {
					fmt.Println("\n\t---Dashboard---")
					fmt.Println("Welcome,", entities.UserData.Name, "!")
					fmt.Println("1. Profile")
					fmt.Println("2. My Own Books")
					fmt.Println("3. See All Books")
					fmt.Println("4. Book I Rented")
					fmt.Print("\n")
					fmt.Print("99. Log Out\n\n")
					fmt.Print("Input: ")
					fmt.Scanln(&inputMenuDashboard)

					switch inputMenuDashboard {
					case 1:
						entities.SeeProfile(conn)
					case 2:
						entities.MyBooks(conn)
					case 3:
						entities.SeeBooks(conn)
					case 4:
						entities.BooksRented()
					case 99:
						fmt.Println("\nExit Dashboard")
					default:
						fmt.Println("\nWrong input menu")
					}
				}
			} else {
				continue
			}
		case 3:
			entities.SeeBooks(conn)
		case 99:
			fmt.Println("\nExiting program...")
		default:
			fmt.Println("\nWrong input Menu")
		}
	}
}
