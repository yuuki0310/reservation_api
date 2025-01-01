package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	UUID      string    `gorm:"type:char(36);unique;not null" json:"sub"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
