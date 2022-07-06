package main

import (
	"RentBookGroupProject/db"
	"RentBookGroupProject/entities/book"
	"RentBookGroupProject/entities/user"
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
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
			Register(conn)
		case 2:
			Login()
		case 3:
			SeeBooks(conn)
		case 99:
			fmt.Println("Exiting program...")
		default:
			fmt.Println("Menu tidak terdaftar")
		}
	}
}

func Register(db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()
	if name == "" || email == "" || password == "" {
		fmt.Println("Name or Email or Password cant be blank!")
		return
	}

	user := user.User{Name: name, Email: email, Password: password}
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("Email Registered")
	} else {
		fmt.Println("Successfully registered")
	}
}

func Login() {}

func SeeBooks(db *gorm.DB) {
	var booksData []book.Book
	result := db.Find(&booksData)

	if result.Error != nil {
		log.Println("Error occured")
	} else {
		for i := 0; i < len(booksData); i++ {
			fmt.Print("\n")
			fmt.Println("ID: ", booksData[i].ID)
			fmt.Println("Name: ", booksData[i].Name)
			fmt.Println("Type: ", booksData[i].Type)
			fmt.Println("Rented? ", booksData[i].Status)
		}
	}
}
