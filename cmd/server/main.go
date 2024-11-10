package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki0310/reservation_api/infrastructure/mysql"
	"github.com/yuuki0310/reservation_api/interfaces"
)

func main() {
	r := gin.Default()

	mysql.InitDatabase()

	interfaces.DefineRoutes(r)

	r.Run(":8080")
}
