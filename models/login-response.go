package models

type LoginResponse struct {
	JWT        string `json:"token"`
	ClientType string `json:"clientType"`
}
