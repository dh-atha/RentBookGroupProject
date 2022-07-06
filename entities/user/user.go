package user

import (
	"RentBookGroupProject/entities/book"
	"RentBookGroupProject/entities/rent"
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string      `gorm:"not null"`
	Email    string      `gorm:"not null;unique"`
	Password string      `gorm:"not null;"`
	Books    []book.Book `gorm:"foreignKey:UserID"`
	Rents    []rent.Rent `gorm:"foreignKey:UserID"`
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
		fmt.Println("\nCant register, Name or Email or Password cant be blank!")
		return
	}

	user := User{Name: name, Email: email, Password: password}
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("Email registered")
	} else {
		fmt.Println("Successfully registered")
	}
}

var userData User

func Login(db *gorm.DB) bool {
	var check bool

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Login---")
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()

	result := db.Where("email = ? AND password = ?", email, password).First(&userData)
	if result.RowsAffected < 1 {
		log.Println("The email or password is incorrect")
	} else {
		log.Println("Login Success")
		check = true
	}
	return check
}

func SeeProfile(db *gorm.DB) {
	//saat lihat profil ada pilihan yg mengarahkan ke opsi edit, delete atau kembali ke dashboard
	var inputMenu int
	fmt.Println("\n---See Profile---")
	fmt.Println("ID:", userData.ID)
	fmt.Println("Name:", userData.Name)
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
