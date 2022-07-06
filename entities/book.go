package entities

import (
	"fmt"
	"log"

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

func AddBook() {}

func SeeBooks(db *gorm.DB) {
	var booksData []Book
	result := db.Find(&booksData)

	if result.Error != nil {
		log.Println("Error occured")
	} else {
		for i := 0; i < len(booksData); i++ {
			fmt.Println("-------------------------------------")
			fmt.Println("ID\t: ", booksData[i].ID)
			fmt.Println("Name\t: ", booksData[i].Name)
			fmt.Println("Type\t: ", booksData[i].Type)
			fmt.Println("Status\t: ", booksData[i].Status)
			fmt.Println("-------------------------------------")
		}
	}

	// fmt.Println(UserData) // Isinya data user yang lagi login
}

func EditBook() {}

func DeleteBook() {}

func MyBooks() {}
