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
		fmt.Println("Please fill field before continue!")
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
	var listBook Book
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\t---Edit Book---")
	fmt.Print("ID Book: ")
	scanner.Scan()
	bid := scanner.Text()

	result := db.Where("id = ? AND user_id = ?", bid, UserData.ID).Find(&listBook)
	if result.RowsAffected < 1 {
		fmt.Println("\nCan't Edit this book")
	} else {
		name := listBook.Name
		tipe := listBook.Type
		status := listBook.Status
		fmt.Println("Choose what you want to edit:\n1. Name\n2. Type\n3. Status")
		fmt.Print("\nInput: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "1" {
			fmt.Print("New Name: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Println()
		} else if input == "2" {
			fmt.Print("New Type: ")
			scanner.Scan()
			tipe = scanner.Text()
			fmt.Println()
		} else if input == "3" {
			fmt.Print("New Status: ")
			fmt.Scan(&status)
			fmt.Println()
		} else {
			fmt.Println("Wrong input menu")
			EditBook(db)
		}

		listBook.Name = name
		listBook.Type = tipe
		listBook.Status = status
		err := db.Save(&listBook)
		if err.Error != nil {
			fmt.Println("Error occured")
		} else {
			log.Println("Succesfully Updated")
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

	result := db.Where("id = ? AND user_id = ?", bid, UserData.ID).Find(&listBook)
	if result.RowsAffected < 1 {
		fmt.Println("\nThis book is not yours")
	} else {
		var input string
		fmt.Println("Are you sure to delete this book?\nYour book title is ", listBook.Name, " (Y/N)")
		fmt.Scan(&input)
		fmt.Println()
		if input == "Y" || input == "y" {
			err := db.Delete(&listBook)
			if err.Error != nil {
				fmt.Println("\nCan't delete this book")
			}
			log.Println("Your book was deleted")
		} else if input == "N" || input == "n" {
			fmt.Println("Canceled")
		} else {
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
