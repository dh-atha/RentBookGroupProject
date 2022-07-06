package register

import (
	"RentBookGroupProject/entities/user"
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

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
