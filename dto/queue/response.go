package queue

import (
	"antrego/models"
)

type Response struct {
	ID            uint   `json:"id"`
	BookingCode   string `json:"booking_code"`
	ClinicCode    string `json:"clinic_code"`
	Number        string `json:"number"`
	EstimatedTime string `json:"estimated_time"`
	Status        string `json:"status"`
}

func NewResponse(queue models.Queue) Response {
	return Response{
		ID:            queue.ID,
		BookingCode:   queue.BookingCode,
		ClinicCode:    queue.ClinicCode,
		Number:        queue.Number,
		EstimatedTime: queue.EstimatedTime.Format("2006-01-02 15:04"),
		Status:        queue.Status,
	}
}

func NewListResponse(queue []models.Queue) []Response {
	response := make([]Response, len(queue))
	for i, q := range queue {
		response[i] = NewResponse(q)
	}
	return response
}
