package user

import (
	"RentBookGroupProject/auth/login"
	"fmt"

	"gorm.io/gorm"
)

func SeeProfile(db *gorm.DB) {
	//saat lihat profil ada pilihan yg mengarahkan ke opsi edit, delete atau kembali ke dashboard
	fmt.Println("\n---See Profile---")
	userData := login.GetUserData()
	fmt.Println(userData)
}

func EditProfile() {
	//ketika edit profil berhasil langsung kembali ke seeprofile
}

func DeleteProfile() {
	//ketika delete profile berhasil langsung kembali ke homemenu
}
