package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuuki0310/reservation_api/infrastructure"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
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
