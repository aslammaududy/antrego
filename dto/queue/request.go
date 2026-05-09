package queue

type BookRequest struct {
	ClinicCode string `json:"clinic_code" binding:"required"`
}
