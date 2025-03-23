package models

type Train struct {
	Id uint8 `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Image         string `json:"image"`
	Video_url     string `json:"video_url"`
	Difficult     uint8  `json:"difficult"`
	Duration_time int64  `json:"duration_time"`

	Lead_muscle_id uint8 `json:"lead_muscle_id"`
}

type TrainWithMuscle struct {
	Id uint8 `json:"id"`

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Image         string `json:"image,omitempty"`
	Video_url     string `json:"video_url,omitempty"`
	Difficult     uint8  `json:"difficult,omitempty"`
	Duration_time int64  `json:"duration_time,omitempty"`

	Muscles Muscles `json:"muscles"`
}
