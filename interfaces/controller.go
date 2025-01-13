package interfaces

import (
	"fmt"
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
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get store reservations"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func getUser(c *gin.Context) *model.User {
	user := &model.User{}
	uuid := c.Param("uuid")

	user = &model.User{
		UUID:  uuid,
		Email: "test",
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
	}
	return user
}

func createUser(c *gin.Context) {
	user := getUser(c)
	userRepository := infrastructure.NewUserRepository()
	err := userRepository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func createReservations(c *gin.Context) {
	uuid := c.Param("uuid")
	var req struct {
		StoreID int       `json:"store_id" binding:"required"`
		Date    time.Time `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid reservation data"})
		return
	}

	userRepository := infrastructure.NewUserRepository()
	userID, err := userRepository.GetUserIDByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ユーザーが見つかりません"})
		return
	}

	reservation := &model.Reservation{
		UserID:  userID,
		StoreID: uint(req.StoreID),
		Date:    req.Date,
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
			"date":    reservation.Date.Format("2006/01/02(月)"),
		})
	}

	c.JSON(http.StatusOK, response)
}
