package infrastructure

import (
	"time"

	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/domain/repository"
	"github.com/yuuki0310/reservation_api/infrastructure/mysql"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository() repository.ReservationRepository {
	return &reservationRepository{db: mysql.DB}
}

func (r *reservationRepository) GetReservationsByStoreIDAndDateRange(storeID int, startDate, endDate time.Time) (reservations []*model.Reservation, err error) {
	if err := r.db.Where("store_id = ? AND date >= ? AND date <= ?", storeID, startDate, endDate).
		Find(&reservations).Error; err != nil {
		return nil, err
	}
	return
}

func (r *reservationRepository) CreateReservation(reservation *model.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationRepository) GetReservationsByUserID(userID uint) ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	if err := r.db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
