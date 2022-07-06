package users

import (
	books "RentBookGroupProject/entities/book"
	rents "RentBookGroupProject/entities/rent"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string       `gorm:"not null"`
	Email    string       `gorm:"not null;unique"`
	Password string       `gorm:"not null;"`
	Books    []books.Book `gorm:"foreignKey:UserID"`
	Rents    []rents.Rent `gorm:"foreignKey:UserID"`
}
