package rents

import (
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID uint
	BookID uint
}
