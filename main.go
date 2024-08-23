package main

import (
	"context"
	"log"
	"time"

	"grpc-gateway/configs"
	"grpc-gateway/controllers"
	pb "grpc-gateway/pb"
	"grpc-gateway/routers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := configs.GatewayConfig()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.NewClient(config.UserHost+":"+config.Userport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())

	userController := controllers.InitUserController(c, ctx)

	mainGroup := router.Group("/api/v1")
	routers.UserRoutes(mainGroup, userController)
	router.Run(":" + config.Serviceport)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
