package book

import (
	"RentBookGroupProject/entities/rent"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Type   string `gorm:"not null"`
	Status bool   `gorm:"type:bool;default:false"`
	UserID uint
	Rents  []rent.Rent `gorm:"foreignKey:BookID"`
}
