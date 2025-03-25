package models

type Achievenment struct {
	Id          uint8  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descripton"`
	Image       string `json:"image"`
}
