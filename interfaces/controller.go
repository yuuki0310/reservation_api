package interfaces

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/infrastructure"
)

var (
    availableTimeSlots = []struct {
        From time.Time
        To   time.Time
    }{
        {From: time.Date(0, 1, 1, 9, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 11, 0, 0, 0, time.Local)},
        {From: time.Date(0, 1, 1, 11, 30, 0, 0, time.Local), To: time.Date(0, 1, 1, 13, 30, 0, 0, time.Local)},
        {From: time.Date(0, 1, 1, 14, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 16, 0, 0, 0, time.Local)},
        {From: time.Date(0, 1, 1, 16, 30, 0, 0, time.Local), To: time.Date(0, 1, 1, 18, 30, 0, 0, time.Local)},
        {From: time.Date(0, 1, 1, 19, 0, 0, 0, time.Local), To: time.Date(0, 1, 1, 21, 0, 0, 0, time.Local)},
        {From: time.Date(0, 1, 1, 21, 30, 0, 0, time.Local), To: time.Date(0, 1, 1, 23, 0, 0, 0, time.Local)},
    }
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createReservation(c *gin.Context) {
    var req struct {
        UUID    string    `json:"uuid" binding:"required"`
        StoreID int       `json:"store_id" binding:"required"`
        From    time.Time `json:"from" binding:"required"`
        To      time.Time `json:"to" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid reservation data"})
        return
    }

	// 予約可能時間帯のチェック
    isValidTime := false
    for _, slot := range availableTimeSlots {
        if req.From.Hour() == slot.From.Hour() && req.From.Minute() == slot.From.Minute() &&
            req.To.Hour() == slot.To.Hour() && req.To.Minute() == slot.To.Minute() {
            isValidTime = true
            break
        }
    }

    if !isValidTime {
        c.JSON(http.StatusBadRequest, gin.H{"message": "予約時間が無効です"})
        return
    }

    userRepository := infrastructure.NewUserRepository()
    userID, err := userRepository.GetUserIDByUUID(req.UUID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "ユーザーが見つかりません"})
        return
    }

	reservation := &model.Reservation{
        UserID:      userID,
        StoreID:     uint(req.StoreID),
        FromDatetime: req.From,
        ToDatetime:   req.To,
    }

    reservationRepository := infrastructure.NewReservationRepository()
    err = reservationRepository.CreateReservation(reservation)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "予約の作成に失敗しました"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully"})
}

func userReservations(c *gin.Context) {
    uuid := c.Param("uuid")

    userRepository := infrastructure.NewUserRepository()
    userID, err := userRepository.GetUserIDByUUID(uuid)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No reservations found for this user"})
        return
    }

    reservationRepository := infrastructure.NewReservationRepository()
    reservations, err := reservationRepository.GetReservationsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No reservations found for this user"})
        return
    }

    var response []gin.H
    for _, reservation := range reservations {
        response = append(response, gin.H{
            "storeId": reservation.StoreID,
            "date":    reservation.FromDatetime.Format("2006/01/02(月)"),
            "from":    reservation.FromDatetime.Format("15:04"),
            "to":      reservation.ToDatetime.Format("15:04"),
        })
    }

    c.JSON(http.StatusOK, response)
}
