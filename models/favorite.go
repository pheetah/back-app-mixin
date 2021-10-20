package models

type Favorite struct {
	ID      int    `json:"id"`
	Favtype string `json:"favType"`
	Alias   string `json:"alias"`
	Date    string `json:"date"`
}
