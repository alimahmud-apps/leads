package models

import (
	"time"
)

type Lead struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Source      string    `json:"source" db:"source"`
	Message     string    `json:"message" db:"message"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type GeneralErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type GeneralSuccessResponse struct {
	Status string                 `json:"status"`
	ID     int                    `json:"id,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

type RequestSubmitLead struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone" validate:"required"`
	Source      string `json:"source"`
	Message     string `json:"message"`
}

type ErrorLogs struct {
	ID           int       `json:"id" db:"id"`
	ErrorMessage string    `json:"error_message" db:"error_message"`
	Endpoint     string    `json:"endpoint" db:"endpoint"`
	StatusCode   string    `json:"status_code" db:"status_code"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
}
