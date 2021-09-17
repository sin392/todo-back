package model

import "time"

type (
	TODO struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	ReadTODORequest struct {
		ID int `json:"id"`
	}
	ReadTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	UpdateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	UpdateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	DeleteTODORequest  struct{}
	DeleteTODOResponse struct{}
)
