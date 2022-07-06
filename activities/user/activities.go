package user

import (
	"fmt"

	"gorm.io/gorm"
)

func SeeProfile(db *gorm.DB) {
	//saat lihat profil ada pilihan yg mengarahkan ke opsi edit, delete atau kembali ke dashboard
	// var inputMenu int
	// userData := login.GetUserData()
	// fmt.Println("\n---See Profile---")
	// fmt.Println("ID:", userData.ID)
	// fmt.Println("Name:", userData.Name)
	// fmt.Println("1. Edit")
	// fmt.Println("2. Delete")
	// fmt.Print("\nInput: ")
	// fmt.Scanln(&inputMenu)
	// switch inputMenu {
	// case 1:
	// 	EditProfile()
	// case 2:
	// 	DeleteProfile()
	// }
}

func EditProfile() {
	//ketika edit profil berhasil langsung kembali ke seeprofile
	fmt.Println("EDIT")
}

func DeleteProfile() {
	//ketika delete profile berhasil langsung kembali ke homemenu
	fmt.Println("DELETE")
}
