package interfaces

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	route := r.Group("/")

	route.GET("/", test)

	stores := route.Group("/stores")
	{
		stores.GET("/:storeId/reservations", storeReservations)
	}

	users := route.Group("/users")
	users.Use(jwtMiddleware())
	{
		users.POST("", createUser)
		users.POST("/:uuid/reservations", createReservations)
		users.GET("/:uuid/reservations", userReservations)
	}
}
