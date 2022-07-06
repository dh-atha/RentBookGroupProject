package main

import (
	"RentBookGroupProject/entities/book"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SeeProfile(db *gorm.DB) {
	//saat lihat profil ada pilihan yg mengarahkan ke opsi edit, delete atau kembali ke dashboard
	fmt.Println("\n---See Profile---")
	userData := GetUserData()
	fmt.Println("Name     : ", userData.Name)
	fmt.Println("Email    : ", userData.Email)
	fmt.Println("Password : ", userData.Password)
}

func EditProfile() {
	//ketika edit profil berhasil langsung kembali ke seeprofile
}

func DeleteProfile() {
	//ketika delete profile berhasil langsung kembali ke homemenu
}

func SeeBooks(db *gorm.DB) {
	var booksData []book.Book
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
}

func EditBook() {}

func DeleteBook() {}

func MyBooks() {}

func RentBook() {}

func BooksRented() {}

func ReturnBook() {}
