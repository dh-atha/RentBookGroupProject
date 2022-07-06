package entities

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Type   string `gorm:"not null"`
	Status bool   `gorm:"type:bool;default:false"`
	UserID uint
	Rents  []Rent `gorm:"foreignKey:BookID"`
}

func AddBook(db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Add Book---")
	fmt.Print("Title: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Type: ")
	scanner.Scan()
	tipe := scanner.Text()
	fmt.Println()

	fmt.Println()
	if name == "" || tipe == "" {
		fmt.Println("Please fill the blank!")
		return
	}

	book := Book{Name: name, Type: tipe, Status: true, UserID: UserData.ID}
	result := db.Create(&book)

	if result.Error != nil {
		fmt.Println("Can't add your book.")
	} else {
		log.Println("Succesfully add your book")
	}
	MyBooks(db)
}

func SeeBooks(db *gorm.DB) {
	var booksData []Book
	var inputMenu int
	result := db.Find(&booksData)

	if result.Error != nil {
		log.Println("Error occured")
	} else {
		fmt.Println("\n\t---Book List---")
		for i := 0; i < len(booksData); i++ {
			var userName User
			fmt.Println("-------------------------------------")
			fmt.Println("ID\t: ", booksData[i].ID)
			fmt.Println("Name\t: ", booksData[i].Name)
			fmt.Println("Type\t: ", booksData[i].Type)
			db.Where("id = ?", booksData[i].UserID).Find(&userName)
			fmt.Println("Owner\t: ", userName.Name)
			if booksData[i].Status == true {
				fmt.Println("Status\t:  Available")
			} else {
				fmt.Println("Status\t:  Not available")
			}
			if i == len(booksData)-1 {
				fmt.Println("-------------------------------------")
			}
		}
	}
	fmt.Println("\n00. Previous")
	fmt.Print("\nInput: ")
	fmt.Scanln(&inputMenu)
	switch inputMenu {
	case 00:
		return
	default:
		fmt.Println("\nWrong input menu")
		SeeBooks(db)
	}

}

func EditBook(db *gorm.DB) {
	var input int
	var listBook Book
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Edit Book---")
	fmt.Print("ID Book: ")
	scanner.Scan()
	bid := scanner.Text()
	fmt.Println("Choose what you want to edit:\n1. Name\n2. Type\n3. Status")
	fmt.Print("\nInput: ")
	fmt.Scan(&input)
	fmt.Println()

	result := db.Where("id = ? AND user_id = ?", bid, UserData.ID).Find(&listBook)
	if result.RowsAffected < 1 {
		fmt.Println("Can't Edit this book")
	} else {
		if input == 1 {
			fmt.Print("\nNew Name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println()

			err := db.Model(&listBook).Update("name", name)
			if err.Error != nil {
				fmt.Println("Error occured")
			}
			log.Println("Succesfully Updated")
		} else if input == 2 {
			fmt.Print("\nNew Type: ")
			scanner.Scan()
			tipe := scanner.Text()
			fmt.Println()

			err := db.Model(&listBook).Update("type", tipe)
			if err.Error != nil {
				fmt.Println("Error occured")
			}
			log.Println("Succesfully Updated")
		} else if input == 3 {
			var status bool
			fmt.Print("\nNew Status: ")
			fmt.Scan(&status)
			fmt.Println()

			err := db.Model(&listBook).Update("status", status)
			if err.Error != nil {
				fmt.Println("Error occured")
			}
			log.Println("Succesfully Updated")
		} else {
			fmt.Println("Wrong input menu")
			EditBook(db)
		}
	}
	MyBooks(db)

}

func DeleteBook(db *gorm.DB) {
	var listBook Book
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Delete Book---")
	fmt.Print("ID Book: ")
	scanner.Scan()
	bid := scanner.Text()
	fmt.Println()

	result := db.Where("id = ? AND user_id = ?", bid, UserData.ID).Find(&listBook)
	if result.RowsAffected < 1 {
		fmt.Println("This book is not yours")
	} else {
		var input string
		fmt.Print("Are you sure to delete this book?\nYour book title is ", listBook.Name, "(Y/N)")
		fmt.Scan(&input)
		fmt.Println()
		switch input {
		case "Y":
			err := db.Delete(&listBook)
			if err.Error != nil {
				fmt.Println("Can't delete this book")
			}
			log.Println("Your book was deleted")
		case "N":
			fmt.Println("Canceled")
		default:
			fmt.Println("Wrong input menu")
			DeleteBook(db)
		}
	}
	MyBooks(db)
}

func MyBooks(db *gorm.DB) {
	var listBook []Book
	var inputMenu int
	result := db.Where("user_id = ?", UserData.ID).Find(&listBook)

	if result.Error != nil {
		log.Println("Error occured")
	} else {
		fmt.Println("\n\t---Book List---")
		for i := 0; i < len(listBook); i++ {
			fmt.Println("-------------------------------------")
			fmt.Println("ID\t: ", listBook[i].ID)
			fmt.Println("Name\t: ", listBook[i].Name)
			fmt.Println("Type\t: ", listBook[i].Type)
			if listBook[i].Status == true {
				fmt.Println("Status\t:  Available")
			} else {
				fmt.Println("Status\t:  Not available")
			}
			if i == len(listBook)-1 {
				fmt.Println("-------------------------------------")
			}
		}

	}
	fmt.Println("\n1. Add Book")
	fmt.Println("2. Edit Book")
	fmt.Println("3. Delete Book")
	fmt.Println("00. Previous")
	fmt.Print("\nInput: ")
	fmt.Scanln(&inputMenu)
	switch inputMenu {
	case 1:
		AddBook(db)
	case 2:
		EditBook(db)
	case 3:
		DeleteBook(db)
	case 00:
		return
	default:
		fmt.Println("\nWrong input menu")
		MyBooks(db)
	}
}
