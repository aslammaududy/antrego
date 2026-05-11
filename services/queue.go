package services

import (
	"antrego/config"
	"antrego/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateQueue(c *gin.Context, clinicCode string) (models.Queue, error) {
	queues, err := gorm.G[models.Queue](config.DB).Where("clinic_code = ?", clinicCode).Find(c.Request.Context())
	if err != nil {
		return models.Queue{}, err
	}

	queue := models.Queue{
		BookingCode:   generateBookingCode(len(queues) + 1),
		ClinicCode:    clinicCode,
		Number:        len(queues) + 1,
		EstimatedTime: estimate(len(queues)),
		Status:        "booked",
	}
	result := gorm.WithResult()
	err = gorm.G[models.Queue](config.DB, result).Create(c.Request.Context(), &queue)
	if err != nil {
		return models.Queue{}, err
	}
	return queue, nil
}

func generateBookingCode(sequence int) string {
	return fmt.Sprintf("%s%03d", time.Now().Format("20060102"), sequence)
}

func estimate(totalQueue int) time.Time {
	startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local)
	waitingTimePerPatient := 5

	waitingTime := totalQueue * waitingTimePerPatient

	return startTime.Add(time.Duration(waitingTime))
}

func UpdateQueueStatus(c *gin.Context, bookingCode, status string) (models.Queue, error) {
	_, err := gorm.G[models.Queue](config.DB).Where("booking_code = ?", bookingCode).Update(c.Request.Context(), "status", status)
	if err != nil {
		return models.Queue{}, err
	}

	myQueue, err := gorm.G[models.Queue](config.DB).Where("booking_code = ?", bookingCode).First(c.Request.Context())
	if err != nil {
		return models.Queue{}, err
	}

	return myQueue, nil
}
