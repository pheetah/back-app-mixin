package models

type ChangeUserType struct {
	ClientType ClientType `json:"clientType"`
	Email      EmailType  `json:"email"`
}

type ChangeUserTypeBody struct {
	ClientType ClientType `json:"clientType"`
}
