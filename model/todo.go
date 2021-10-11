package model

import "time"

type (
	// Request
	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	ReadTODORequest struct {
		ID string `json:"id"`
	}
	ReadTODOsRequest struct {
		Skip  *string `json:"skip"`
		Limit *string `json:"limit"`
	}
	UpdateTODORequest struct {
		ID          string `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	DeleteTODORequest struct {
		ID string `json:"id"`
	}
	DeleteTODOsRequest struct {
		IDs []string `json:"ids"`
	}
	// Response
	TODO struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	TODOResponse struct {
		TODO TODO `json:"todo"`
	}
	TODOsResponse struct {
		TODOs []TODO `json:"todos"`
	}
)
