package dto

import (
	"encoding/json"
	"time"
)

type ErrorDTO struct {
	Timestamp string `json:"timestamp"`
	Detail    string `json:"detail"`
}

func NewErrorDTO(detail string) *ErrorDTO {
	return &ErrorDTO{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Detail:    detail,
	}
}

func (e *ErrorDTO) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}
