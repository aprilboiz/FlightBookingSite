package dto

type ErrorResponse struct {
	Status  int         `json:"status"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
