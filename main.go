package main

import (
	"RentBookGroupProject/db"
	"RentBookGroupProject/entities/user"
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

var userData []user.User
var InputMenu int

func main() {
	conn := db.InitDB()
	db.MigrateDB(conn)
	i := 1
	for InputMenu != 99 {
		if i == 1 {
			HomeMenu(conn)
		} else {
			DashboardMenu(conn)
		}
		i++
	}
}

func HomeMenu(db *gorm.DB) {
	fmt.Println("---Home Menu---")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. See Books")
	fmt.Print("\n")
	fmt.Print("99. Exit\n\n")
	fmt.Print("Input: ")
	fmt.Scanln(&InputMenu)

	switch InputMenu {
	case 1:
		Register(db)
	case 2:
		Login(db)
		DashboardMenu(db)
	case 3:
		SeeBooks(db)
	case 99:
		fmt.Println("Exiting program...")
		break
	default:
		fmt.Println("Menu tidak terdaftar")
	}
}

func DashboardMenu(db *gorm.DB) {
	fmt.Println("\n---Dashboard---")
	fmt.Println("1. Profile")
	fmt.Println("2. My Own Books")
	fmt.Println("3. Book I Rented")
	fmt.Println("4. Return Books")
	fmt.Print("\n")
	fmt.Print("99. Exit\n\n")
	fmt.Print("Input: ")
	fmt.Scanln(&InputMenu)

	switch InputMenu {
	case 1:
		SeeProfile(db)
	case 2:
		MyBooks()
	case 3:
		BooksRented()
	case 4:
		ReturnBook()

	case 99:
		fmt.Println("Exiting program...")
		break
	default:
		fmt.Println("Menu tidak terdaftar")
	}
}

func Register(db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Register---")
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
		Register(db)
		return
	}

	user := user.User{Name: name, Email: email, Password: password}
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("Email registered")
	} else {
		fmt.Println("Successfully registered")
	}
}

func Login(db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Login---")
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()

	result := db.Where("email = ? AND password = ?", email, password).Find(&userData)
	if result.RowsAffected < 1 {
		log.Println("The email or password is incorrect")
		Login(db)
	} else {
		log.Println("Login Success")
	}
}

func GetUserData() user.User {
	return userData[0]
}
