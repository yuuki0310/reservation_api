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

type StoreReservation struct {
	StoreID      int                `json:"storeId"`
	Year         int                `json:"year"`
	Month        int                `json:"month"`
	Reservations []DailyReservation `json:"reservations"`
}

type DailyReservation struct {
	Date      time.Time         `json:"date"`
	Weekday   int               `json:"weekday"`   // 日曜日を0とした曜日
	IsHoliday bool              `json:"isHoliday"` // 祝日かどうか
	Slots     []ReservationSlot `json:"slots"`
}

type ReservationSlot struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Status string    `json:"status"` // 予約ステータス ("available" または "booked")
}

var AvailableTimeSlots = []ReservationSlot{
	{From: time.Date(0, 1, 1, 0, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 1, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 2, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 3, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 4, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 5, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 6, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 7, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 8, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 9, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 10, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 11, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 12, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 13, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 14, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 15, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 16, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 17, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 18, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 19, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 20, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 21, 45, 0, 0, time.Local)},
	{From: time.Date(0, 1, 1, 22, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 23, 45, 0, 0, time.Local)},
}
