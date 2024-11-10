package interfaces

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	v1 := r.Group("/")
	{
		v1.GET("/", test)
	}
	{
		 v1.POST("/reservations", createReservation)
		v1.GET("/users/:uuid/reservations", userReservations)
	}
}
