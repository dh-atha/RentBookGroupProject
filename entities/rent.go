package entities

import (
	"bufio"
	"fmt"
	"log"
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
	for inputMenu != 99 {
		var listRents []Rent
		result := db.Where("user_id = ?", UserData.ID).Find(&listRents)

		if result.Error != nil {
			log.Println("Error occured")
		} else {
			fmt.Println("\n\t---Rented Book List---")
			for i := 0; i < len(listRents); i++ {
				bookData := Book{}
				db.Where("id = ?", listRents[i].BookID).Find(&bookData)
				fmt.Println("-------------------------------------")
				fmt.Println("Rent ID\t: ", listRents[i].ID)
				fmt.Println("Book ID\t: ", listRents[i].BookID)
				fmt.Println("Book Title\t: ", bookData.Name)
				fmt.Println("Rent Date\t: ", listRents[i].CreatedAt)

				if i == len(listRents)-1 {
					fmt.Println("-------------------------------------")
				}
			}

		}
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
	var rentData Rent

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n---Return Book---")
	fmt.Print("RentID: ")
	scanner.Scan()
	rentID := scanner.Text()

	res := db.Where("id = ? AND user_id = ?", rentID, UserData.ID).Find(&rentData)

	if res.Error != nil {
		log.Println("Error Occured")
	} else {
		if res.RowsAffected < 1 {
			fmt.Println("Cant delete that rent data, its either not yours or wrong rentID input")
		} else {
			fmt.Println("Successfully returning book")
			var bookData Book
			db.Where("id = ?", rentData.BookID).Find(&bookData)
			bookData.Status = true
			db.Save(&bookData)
			db.Delete(&rentData)
		}
	}
}
