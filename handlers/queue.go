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

// GetAllQueues
// @Summary List all queues
// @Description Get all queue records
// @Tags queues
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{data=[]queue.Response}
// @Router /queues [get]
func GetAllQueue(c *gin.Context) {
	queues, err := gorm.G[models.Queue](config.DB).Find(c.Request.Context())
	if err != nil {
		c.Error(errors.New(err.Error()))
	}

	dto.OK(c, queue.NewListResponse(queues))
}

// BookTheQueue
// @Summary Book a queue
// @Description Create a new queue booking
// @Tags queues
// @Accept json
// @Produce json
// @Param request body queue.BookRequest true "Booking details"
// @Success 200 {object} dto.Response{data=queue.Response}
// @Router /queue [post]
func BookTheQueue(c *gin.Context) {
	var request queue.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(middleware.ErrBadRequest)
		return
	}

	myQueue, err := services.GenerateQueue(c, request.ClinicCode)
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	dto.OK(c, queue.NewResponse(myQueue))
}

// MyQueue
// @Summary Get a queue by booking code
// @Description Get queue details by booking code
// @Tags queues
// @Accept json
// @Produce json
// @Param bookingCode path string true "Booking code"
// @Success 200 {object} dto.Response{data=queue.Response}
// @Failure 404 {object} dto.Response
// @Router /queue/{bookingCode} [get]
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

// CallTheQueue
// @Summary Call a queue
// @Description Mark a queue as called
// @Tags queues
// @Accept json
// @Produce json
// @Param bookingCode path string true "Booking code"
// @Success 200 {object} dto.Response{data=queue.Response}
// @Router /queue/{bookingCode}/call [patch]
func CallTheQueue(c *gin.Context) {
	var bookingCode = c.Param("bookingCode")

	myQueue, err := services.UpdateQueueStatus(c, bookingCode, "called")
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	dto.OK(c, queue.NewResponse(myQueue))
}

// CompleteTheQueue
// @Summary Complete a queue
// @Description Mark a queue as done
// @Tags queues
// @Accept json
// @Produce json
// @Param bookingCode path string true "Booking code"
// @Success 200 {object} dto.Response{data=queue.Response}
// @Router /queue/{bookingCode}/complete [patch]
func CompleteTheQueue(c *gin.Context) {
	var bookingCode = c.Param("bookingCode")

	myQueue, err := services.UpdateQueueStatus(c, bookingCode, "done")
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	dto.OK(c, queue.NewResponse(myQueue))
}

// CancelTheQueue
// @Summary Cancel a queue
// @Description Cancel a queue booking (cannot cancel if status is done)
// @Tags queues
// @Accept json
// @Produce json
// @Param bookingCode path string true "Booking code"
// @Success 200 {object} dto.Response{data=queue.Response}
// @Failure 422 {object} dto.Response
// @Router /queue/{bookingCode}/cancel [patch]
func CancelTheQueue(c *gin.Context) {
	var bookingCode = c.Param("bookingCode")

	lastQueueStatus, err := gorm.G[models.Queue](config.DB).Where("booking_code = ?", bookingCode).Select("status").First(c.Request.Context())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(middleware.ErrNotFound)
			return
		} else {
			c.Error(errors.New(err.Error()))
			return
		}
	}

	if lastQueueStatus.Status == "done" {
		c.Error(middleware.ErrInvalidStatus)
		return
	}

	myQueue, err := services.UpdateQueueStatus(c, bookingCode, "cancelled")
	if err != nil {
		c.Error(errors.New(err.Error()))
		return
	}

	dto.OK(c, queue.NewResponse(myQueue))
}
