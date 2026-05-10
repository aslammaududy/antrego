package services

import (
	"antrego/config"
	"antrego/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GenerateQueue(clinicCode string) error {
	queues, err := gorm.G[models.Queue](config.DB).Where("clinic_code = ?", clinicCode).Find(config.Ctx)
	if err != nil {
		return err
	}

	queue := models.Queue{
		BookingCode:   generateBookingCode(len(queues) + 1),
		ClinicCode:    clinicCode,
		Number:        len(queues) + 1,
		EstimatedTime: estimate(len(queues)),
	}

	err = gorm.G[models.Queue](config.DB).Create(config.Ctx, &queue)
	if err != nil {
		return err
	}
	return nil
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
