package handlers

import (
	"antrego/config"
	"antrego/dto/queue"
	"antrego/models"
	"antrego/services"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllQueue(c *gin.Context) {
	queues, err := gorm.G[models.Queue](config.DB).Find(c.Request.Context())
	if err != nil {
		c.Error(errors.New(err.Error()))
	}
	c.JSON(200, queue.NewListResponse(queues))
}

func BookTheQueue(c *gin.Context) {
	var request queue.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	err := services.GenerateQueue(c, request.ClinicCode)
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}
	c.JSON(201, gin.H{"success": true})
}

func MyQueue(c *gin.Context) {
	var bookingCode = c.Param("bookingCode")

	myQueue, err := gorm.G[models.Queue](config.DB).Where("booking_code = ?", bookingCode).First(c.Request.Context())
	if err != nil {
		c.Error(errors.New(err.Error()))
	}

	c.JSON(200, queue.NewResponse(myQueue))
}
