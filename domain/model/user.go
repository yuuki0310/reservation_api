package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	UUID      string    `gorm:"type:char(36);unique;not null"`
	Username  string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
