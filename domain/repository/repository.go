package repository

import (
	"github.com/yuuki0310/reservation_api/domain/model"
)

type ReservationRepository interface {
    GetReservationsByUserID(uint) ([]*model.Reservation, error)
	CreateReservation(*model.Reservation) error
}

type UserRepository interface {
	GetUserIDByUUID(string) (uint, error)
}