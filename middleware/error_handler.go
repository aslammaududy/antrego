package middleware

import (
	"antrego/dto"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound      = &AppError{Status: 404, Code: "NOT_FOUND", Message: "resource not found"}
	ErrUnauthorized  = &AppError{Status: 401, Code: "UNAUTHORIZED", Message: "authentication required"}
	ErrBadRequest    = &AppError{Status: 400, Code: "BAD_REQUEST", Message: "invalid request"}
	ErrInvalidStatus = &AppError{Status: 422, Code: "INVALID_STATUS", Message: "can't cancel the queue. Queue status is done"}
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request first

		// Check if any errors were added to the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := errors.AsType[*AppError](err); ok {
				dto.Fail(c, appErr.Status, appErr.Code, appErr.Message)

			} else {
				dto.Fail(c, http.StatusInternalServerError, "INTERNAL", err.Error())
			}
		}
	}
}
