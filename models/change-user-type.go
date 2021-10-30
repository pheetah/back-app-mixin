package models

type ChangeUserType struct {
	ClientType ClientType `json:"clientType"`
	Email      EmailType  `json:"email"`
}

type ChangeUserTypeRequestBody struct {
	ClientType ClientType `json:"clientType"`
}
