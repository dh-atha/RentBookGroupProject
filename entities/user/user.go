package user

import (
	"RentBookGroupProject/entities/book"
	"RentBookGroupProject/entities/rent"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string      `gorm:"not null"`
	Email    string      `gorm:"not null;unique"`
	Password string      `gorm:"not null;"`
	Books    []book.Book `gorm:"foreignKey:UserID"`
	Rents    []rent.Rent `gorm:"foreignKey:UserID"`
}
