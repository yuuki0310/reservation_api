package model

import "time"

type Reservation struct {
	ID           uint      `gorm:"primary_key;auto_increment"`
	UserID       uint      `gorm:"not null"`
	StoreID      uint      `gorm:"not null"`
	FromDatetime time.Time `gorm:"type:datetime;not null"`
	ToDatetime   time.Time `gorm:"type:datetime;not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
