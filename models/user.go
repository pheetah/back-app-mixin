package models

type ClientType string
type EmailType string

type User struct {
	ID         int        `json:"id"`
	Email      EmailType  `json:"email"`
	Password   string     `json:"password"`
	ClientType ClientType `json:"clientType"`
}
