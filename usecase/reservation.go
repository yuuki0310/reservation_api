package usecase

// usecase層はスキップする

// import (
// 	"github.com/yuuki0310/reservation_api/domain/model"
// 	"github.com/yuuki0310/reservation_api/domain/repository"
// )

// type ReservationUseCase interface {
// 	GetReservationsByUserID(uint) ([]*model.Reservation, error)
// }

// type reservationUseCase struct {
// 	reservationRepository repository.ReservationRepository
// }

// func NewReservationUseCase(r repository.ReservationRepository) ReservationUseCase {
// 	return &reservationUseCase{
// 		reservationRepository: r,
// 	}
// }

// func (u reservationUseCase) GetReservationsByUserID(userID uint) (reservations []*model.Reservation,  err error) {
// 	reservations ,err = u.reservationRepository.GetReservationsByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }