package usecase

import (
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/domain/repository"
	"github.com/yuuki0310/reservation_api/utils"
)

type ReservationUseCase interface {
	GetStoreReservations(int, int, int) (*model.StoreReservation, error)
}

type reservationUseCase struct {
	reservationRepository repository.ReservationRepository
}

func NewReservationUseCase(r repository.ReservationRepository) ReservationUseCase {
	return &reservationUseCase{
		reservationRepository: r,
	}
}

func (u reservationUseCase) GetStoreReservations(storeID, year, month int) (*model.StoreReservation, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, utils.JST)
	endDate := startDate.AddDate(0, 1, -1)

	reservations, err := u.reservationRepository.GetReservationsByStoreIDAndDateRange(storeID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	reservationMap := make(map[time.Time]struct{}, len(reservations))
	for _, reservation := range reservations {
		// DBにJSTで入っているが、UTCで取り出されるため、9時間引く
		reservationMap[reservation.FromDatetime.In(utils.JST).Add(-9*time.Hour)] = struct{}{}
	}

	storeReservations := &model.StoreReservation{
		StoreID: storeID,
		Year:    year,
		Month:   month,
	}
	var dailyReservation []model.DailyReservation
	for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
		var reservationSlot []model.ReservationSlot
		for _, slot := range model.AvailableTimeSlots {
			from := time.Date(date.Year(), date.Month(), date.Day(), slot.From.Hour(), slot.From.Minute(), 0, 0, utils.JST)
			to := time.Date(date.Year(), date.Month(), date.Day(), slot.To.Hour(), slot.To.Minute(), 0, 0, utils.JST)
			status := "available"
			if _, ok := reservationMap[from]; ok {
				status = "booked"
			}
			reservationSlot = append(reservationSlot, model.ReservationSlot{
				From:   from,
				To:     to,
				Status: status,
			})
		}
		dailyReservation = append(dailyReservation, model.DailyReservation{
			Date:      date,
			Weekday:   int(date.Weekday()),
			IsHoliday: holiday.IsHoliday(date),
			Slots:     reservationSlot,
		})
	}
	storeReservations.Reservations = dailyReservation

	return storeReservations, nil
}
