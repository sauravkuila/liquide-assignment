package dto

import e "liquide-assignment/pkg/errors"

type CommonResponse struct {
	Status  bool      `json:"success"`
	Errors  []e.Error `json:"errors,omitempty"`
	Message string    `json:"message,omitempty"`
}
