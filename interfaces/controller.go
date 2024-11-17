package interfaces

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/infrastructure"
	"github.com/yuuki0310/reservation_api/usecase"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func storeReservations(c *gin.Context) {
	storeIDStr := c.Param("storeId")
	if storeIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "storeId is required"})
		return
	}
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "storeId must be a valid number"})
		return
	}

	year, err := strconv.Atoi(c.Query("year"))
	if err != nil || year < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Valid year is required"})
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Valid month (1-12) is required"})
		return
	}

	reservationRepository := infrastructure.NewReservationRepository()
	useCaseRepository := usecase.NewReservationUseCase(reservationRepository)
	response, err := useCaseRepository.GetStoreReservations(storeID, year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get store reservations"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func createReservations(c *gin.Context) {
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
	for _, slot := range model.AvailableTimeSlots {
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
		UserID:       userID,
		StoreID:      uint(req.StoreID),
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
