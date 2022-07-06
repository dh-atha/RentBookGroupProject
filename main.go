package main

import (
	"RentBookGroupProject/auth/home"
)

func main() {
	home.HomeMenu()
}

// func SeeBooks(db *gorm.DB) {
// 	var booksData []book.Book
// 	result := db.Find(&booksData)

// 	if result.Error != nil {
// 		log.Println("Error occured")
// 	} else {
// 		for i := 0; i < len(booksData); i++ {
// 			fmt.Print("-------------------------------------")
// 			fmt.Println("ID\t: ", booksData[i].ID)
// 			fmt.Println("Name\t: ", booksData[i].Name)
// 			fmt.Println("Type\t: ", booksData[i].Type)
// 			fmt.Println("Status\t: ", booksData[i].Status)
// 			fmt.Print("-------------------------------------")
// 		}
// 	}
// }
