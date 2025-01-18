package model

import "time"

type Reservation struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	UserID    uint      `gorm:"not null;index"`
	StoreID   uint      `gorm:"not null;index"`
	Date      string    `gorm:"type:date;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type StoreReservation struct {
	StoreID      int                `json:"storeId"`
	Year         int                `json:"year"`
	Month        int                `json:"month"`
	Reservations []DailyReservation `json:"reservations"`
}

type DailyReservation struct {
	Date      string `json:"date"`      // "YYYY-MM-DD"
	Weekday   int    `json:"weekday"`   // 日曜日を0とした曜日
	IsHoliday bool   `json:"isHoliday"` // 祝日かどうか
	Status    string `json:"status"`    // 予約ステータス ("available" または "booked")
}
