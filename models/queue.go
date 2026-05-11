package models

import (
	"time"

	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model
	BookingCode   string `gorm:"unique"`
	ClinicCode    string `gorm:"unique"`
	Number        int
	EstimatedTime time.Time
	Status        string
}
