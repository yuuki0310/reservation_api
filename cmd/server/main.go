// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

func initDatabase() {
    dsn := "root:password@tcp(mysql:3306)/reservation_db?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
}

func main() {
    r := gin.Default()
    initDatabase()

    // Simple ping route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Start the Gin server
    r.Run(":8080")
}
