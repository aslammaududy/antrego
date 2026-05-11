package main

import (
	"antrego/config"
	"antrego/handlers"
	"antrego/middleware"

	"github.com/gin-gonic/gin"
)

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

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	err := r.Run()
	if err != nil {
		return
	}
}
