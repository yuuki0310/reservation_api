package infrastructure

import (
	"github.com/yuuki0310/reservation_api/domain/repository"
	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/infrastructure/mysql"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository() repository.ReservationRepository {
	return &reservationRepository{db: mysql.DB}
}

func (r *reservationRepository)GetReservationsByUserID(userID uint) ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	if err := r.db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
