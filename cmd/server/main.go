package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/yuuki0310/reservation_api/infrastructure/mysql"
	"github.com/yuuki0310/reservation_api/interfaces"
	"github.com/yuuki0310/reservation_api/utils"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")
	mysql.InitDatabase()
	if err := utils.InitTimezone(); err != nil {
		log.Fatalf("Failed to initialize timezone: %v", err)
	}

	r := gin.Default()
	interfaces.DefineRoutes(r)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	// Path をログに出力
	fmt.Printf("Received Path: %s\n", req.Path)

	// リクエスト全体の確認（デバッグ用）
	fmt.Printf("Full Request: %+v\n", req)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("ENV") == "local" {
		r := gin.Default()
		interfaces.DefineRoutes(r)
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	} else {
		lambda.Start(Handler)
	}
}