package models

type ApiError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
