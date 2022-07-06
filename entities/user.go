package entities

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null;"`
	Books    []Book `gorm:"foreignKey:UserID"`
	Rents    []Rent `gorm:"foreignKey:UserID"`
}

func Register(db *gorm.DB) {
	var dataUser User
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Register---")
	fmt.Print("Your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Println()
	if name == "" || email == "" || password == "" {
		fmt.Println("\nCant register, Name or Email or Password cant be blank!")
		return
	}

	user := User{Name: name, Email: email, Password: password}
	result := db.Where("email = ? AND password = ?", email, password).First(&dataUser)

	if result == nil {
		db.Create(&user)
		fmt.Println("Successfully registered")
	} else {
		log.Println("Email registered")
	}
}

var UserData User // Data User yang lagi login

func Login(db *gorm.DB) bool {
	var check bool

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Login---")
	fmt.Print("Your email: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Println()

	result := db.Where("email = ? AND password = ?", email, password).Find(&UserData)
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
	fmt.Println("ID   :", UserData.ID)
	fmt.Println("Name :", UserData.Name)
	fmt.Println("Email:", UserData.Email)
	fmt.Println("\n1. Edit")
	fmt.Println("2. Delete")
	fmt.Println("00. Previous")
	fmt.Print("\nInput: ")
	fmt.Scanln(&inputMenu)
	switch inputMenu {
	case 1:
		EditProfile()
	case 2:
		DeleteProfile()
	case 00:
		return
	default:
		fmt.Println("\nMenu tidak terdaftar")
		SeeProfile(db)
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
