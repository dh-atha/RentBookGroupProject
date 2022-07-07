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
	result := db.Where("email = ? AND password = ?", email, password).Find(&dataUser)

	if result.Error != nil {
		log.Println("Email registered")
	} else {
		db.Create(&user)
		fmt.Println("Successfully registered")
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
		EditProfile(db)
	case 2:
		DeleteProfile(db)
	case 00:
		return
	default:
		fmt.Println("\nMenu tidak terdaftar")
		SeeProfile(db)
	}
}

func EditProfile(db *gorm.DB) {
	//ketika edit profil berhasil langsung kembali ke seeprofile
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Edit Profile---")
	name := UserData.Name
	email := UserData.Email
	pass := UserData.Password
	fmt.Println("Choose what you want to edit:\n1. Name\n2. Email\n3. Password")
	fmt.Print("\nInput: ")
	scanner.Scan()
	input := scanner.Text()
	if input == "1" {
		fmt.Print("New Name: ")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println()
	} else if input == "2" {
		fmt.Print("New Email: ")
		scanner.Scan()
		email = scanner.Text()
		fmt.Println()
	} else if input == "3" {
		fmt.Print("New Password: ")
		scanner.Scan()
		pass = scanner.Text()
		fmt.Println()
	} else {
		fmt.Println("Wrong input menu")
		EditProfile(db)
	}

	UserData.Name = name
	UserData.Email = email
	UserData.Password = pass
	err := db.Save(&UserData)
	if err.Error != nil {
		fmt.Println("Error occured")
	} else {
		log.Println("Succesfully Updated")
	}

}

var InputMenuDashboard int

func DeleteProfile(db *gorm.DB) {
	//ketika delete profile berhasil langsung kembali ke homemenu
	fmt.Println("\n\t---Delete Profile---")
	var input string
	fmt.Println("Are you sure to delete your profile? (Y/N)")
	fmt.Scan(&input)
	fmt.Println()
	if input == "Y" || input == "y" {
		rentData := []Rent{}
		db.Where("User_ID = ?", UserData.ID).Find(&rentData)
		if len(rentData) != 0 {
			fmt.Println("Return book before deleting your account!")
			return
		}
		BookData := Book{}
		err := db.Delete(&UserData)
		db.Where("User_ID = ?", UserData.ID).Delete(&BookData)
		if err.Error != nil {
			fmt.Println("\nCan't delete your profile")
		}
		fmt.Println("Processing...")
		InputMenuDashboard = 99
	} else if input == "N" || input == "n" {
		fmt.Println("Canceled")
	} else {
		fmt.Println("Wrong input menu")
		DeleteProfile(db)
	}
}
