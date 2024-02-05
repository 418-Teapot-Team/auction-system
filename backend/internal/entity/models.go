package entity

type ErrResponse struct {
	Message string `json:"message" example:"invalid value in Authorization header"`
}
