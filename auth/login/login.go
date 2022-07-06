package login

import (
	"RentBookGroupProject/entities/user"
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

func Login(db *gorm.DB) {
	var userData []user.User
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
