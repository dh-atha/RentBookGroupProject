package rent

import "time"

type Rent struct {
	ID        uint `gorm:"primaryKey; not null"`
	UserID    uint
	BookID    uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time `gorm:"autoCreateTime"`
}
