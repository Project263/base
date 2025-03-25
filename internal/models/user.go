package models

type User struct {
	ID              uint   `json:"id"`
	Login           string `json:"login"`
	Password        string `json:"-"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar,omitempty"`
	AdvancedVersion bool   `json:"advanced_version"`
	Phone           string `json:"phone,omitempty"`
	IsVerifiedPhone bool   `json:"is_verified_phone"`
	Email           string `json:"email"`
	IsVerifiedMail  bool   `json:"is_verified_mail"`
	Age             int    `json:"age,omitempty"`
	Height          int    `json:"height,omitempty"`
	Weight          int    `json:"weight,omitempty"`
	Sex             string `json:"sex,omitempty"`
	DayStreak       int    `json:"day_streak"`
	IsTrainToday    bool   `json:"is_train_today"`
	Points          int    `json:"points"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}
