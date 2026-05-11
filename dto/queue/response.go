package queue

import "antrego/models"

type Response struct {
	ID          uint   `json:"id"`
	BookingCode string `json:"booking_code"`
	ClinicCode  string `json:"clinic_code"`
	Number      int    `json:"number"`
	Status      string `json:"status"`
}

func NewResponse(queue models.Queue) Response {
	return Response{
		ID:          queue.ID,
		BookingCode: queue.BookingCode,
		ClinicCode:  queue.ClinicCode,
		Number:      queue.Number,
		Status:      queue.Status,
	}
}

func NewListResponse(queue []models.Queue) []Response {
	response := make([]Response, len(queue))
	for i, q := range queue {
		response[i] = NewResponse(q)
	}
	return response
}
