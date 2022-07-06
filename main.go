package main

import (
	"RentBookGroupProject/activities/book"
	"RentBookGroupProject/activities/rent"
	"RentBookGroupProject/db"
	users "RentBookGroupProject/entities/user"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"gorm.io/gorm"
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
			Register(conn)
		case 2:
			Login(conn)
			var inputMenuDashboard int
			fmt.Println("\n---Dashboard---")
			fmt.Println("1. Profile")
			fmt.Println("2. My Own Books")
			fmt.Println("3. Book I Rented")
			fmt.Println("4. Return Books")
			fmt.Print("\n")
			fmt.Print("99. Exit\n\n")
			fmt.Print("Input: ")
			fmt.Scanln(&inputMenuDashboard)

			for inputMenuDashboard != 99 {
				switch inputMenuDashboard {
				case 1:
					SeeProfile(conn)
				case 2:
					book.MyBooks()
				case 3:
					rent.BooksRented()
				case 4:
					rent.ReturnBook()
				case 99:
					break
				default:
					fmt.Println("Menu tidak terdaftar")
				}
			}
		case 3:
			book.SeeBooks(conn)
		case 99:
			fmt.Println("Exiting program...")
		default:
			fmt.Println("Menu tidak terdaftar")
		}
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

var userData []users.User

func Login(db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Login---")
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()

	// result := db.Raw("SELECT email, password FROM users WHERE email = ? && password = ?", email, password).Scan(&userData)
	result := db.Where("email = ? AND password = ?", email, password).Find(&userData)
	if result.RowsAffected < 1 {
		log.Println("The email or password is incorrect")
		Login(db)
	} else {
		log.Println("Login Success")
	}
}

func SeeProfile(db *gorm.DB) {
	//saat lihat profil ada pilihan yg mengarahkan ke opsi edit, delete atau kembali ke dashboard
	var inputMenu int
	fmt.Println("\n---See Profile---")
	// fmt.Println("ID:", userData.ID)
	// fmt.Println("Name:", userData.Name)
	fmt.Println("1. Edit")
	fmt.Println("2. Delete")
	fmt.Print("\nInput: ")
	fmt.Scanln(&inputMenu)
	switch inputMenu {
	case 1:
		EditProfile()
	case 2:
		DeleteProfile()
	}
}

func EditProfile() {
	//ketika edit profil berhasil langsung kembali ke seeprofile
	fmt.Println("EDIT")
}

func DeleteProfile() {
	//ketika delete profile berhasil langsung kembali ke homemenu
	fmt.Println("DELETE")
}
