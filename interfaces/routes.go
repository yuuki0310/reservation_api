package interfaces

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/test", test)
	}
	{
		v1.GET("/stores/:storeId/reservations", storeReservations)
	}
	{
		v1.POST("/reservations", createReservations)
		v1.GET("/users/:uuid/reservations", userReservations)
	}
}
