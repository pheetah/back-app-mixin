package models

type LoginResponse struct {
	JWT        string     `json:"token"`
	ClientType ClientType `json:"clientType"`
}
