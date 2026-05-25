package main

import (
	"antrego/config"
	"antrego/handlers"
	"antrego/middleware"

	_ "antrego/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Antrego API
// @version 1.0
// @description Queue management system API
// @host localhost:8080
// @BasePath /
func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	config.Connect()

	r.GET("/queues", handlers.GetAllQueue)
	r.POST("/queue", handlers.BookTheQueue)
	r.GET("/queue/:bookingCode", handlers.MyQueue)
	r.PATCH("/queue/:bookingCode/call", handlers.CallTheQueue)
	r.PATCH("/queue/:bookingCode/complete", handlers.CompleteTheQueue)
	r.PATCH("/queue/:bookingCode/cancel", handlers.CancelTheQueue)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	err := r.Run()
	if err != nil {
		return
	}
}
