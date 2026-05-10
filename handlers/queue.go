package handlers

import (
	"antrego/config"
	"antrego/dto"
	"antrego/dto/queue"
	"antrego/middleware"
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

	dto.OK(c, queue.NewListResponse(queues))
}

func BookTheQueue(c *gin.Context) {
	var request queue.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(middleware.ErrBadRequest)
		return
	}

	err := services.GenerateQueue(c, request.ClinicCode)
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	dto.Created(c)
}

func MyQueue(c *gin.Context) {
	var bookingCode = c.Param("bookingCode")

	myQueue, err := gorm.G[models.Queue](config.DB).Where("booking_code = ?", bookingCode).First(c.Request.Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(middleware.ErrNotFound)
			return
		} else {
			c.Error(errors.New(err.Error()))
			return
		}
	}

	dto.OK(c, queue.NewResponse(myQueue))
}
