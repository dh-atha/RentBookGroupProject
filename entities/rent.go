package entities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID uint
	BookID uint
}

func BooksRented(db *gorm.DB) {
	var inputMenu int
	// var listRents []Rent
	// result := db.Where("user_id = ?", UserData.ID).Find(&listRents)

	// if result.Error != nil {
	// 	log.Println("Error occured")
	// } else {
	// 	fmt.Println("\n\t---Rented Book List---")
	// 	for i := 0; i < len(listRents); i++ {
	// 		fmt.Println("-------------------------------------")
	// 		fmt.Println("ID\t: ", listRents[i].ID)
	// 		fmt.Println("Name\t: ", listRents[i].Name)
	// 		fmt.Println("Type\t: ", listRents[i].Type)
	// 		if i == len(listRents)-1 {
	// 			fmt.Println("-------------------------------------")
	// 		}
	// 	}

	// }
	for inputMenu != 99 {
		fmt.Println("\n1. Rent A Book")
		fmt.Println("2. Return Book")
		fmt.Println("\n99. Previous")
		fmt.Print("\nInput: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			RentBook(db)
		case 2:
			ReturnBook(db)
		case 99:
			return
		default:
			fmt.Println("\nWrong input menu")
		}
	}
}

func RentBook(db *gorm.DB) {
	var bookData Book
	var rentData Rent

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Rent Book---")
	fmt.Print("BookID to Rent: ")
	scanner.Scan()
	bookIDToRent := scanner.Text()
	fmt.Print("UserID to Rent: ")
	scanner.Scan()
	userIDToRent := scanner.Text()
	fmt.Println()

	if userIDToRent == strconv.Itoa(int(UserData.ID)) {
		fmt.Println("Cant rent book from yourself")
		return
	}

	res := db.Where("id = ? AND user_id = ?", bookIDToRent, userIDToRent).Find(&bookData)

	if res.RowsAffected < 1 {
		fmt.Println("Couldn't find a book that matches that bookID and userID")
	} else if bookData.Status == false {
		fmt.Println("Book not available to rent")
	} else {
		rentData.UserID = UserData.ID
		rentData.BookID = bookData.ID
		db.Create(&rentData)
		fmt.Println("Successfully rent a book!")
		bookData.Status = false
		db.Save(&bookData)
	}
}

func ReturnBook(db *gorm.DB) {
	fmt.Println("Return Book")
}
