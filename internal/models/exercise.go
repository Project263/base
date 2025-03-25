package models

type Exercises struct {
	Id             uint8  `json:"id"`
	Title          string `json:"title"`
	Image          string `json:"image"`
	Description    string `json:"descripton"`
	Video_url      string `json:"video_url"`
	EquipmentId    uint8  `json:"equipment_id"`
	Sets           uint8  `json:"Sets"`
	Reps           uint8  `json:"Reps"`
	Difficult      uint8  `json:"Difficult"`
	Lead_muscle_id uint8  `json:"lead_muscle_id"`
}

type FullExercises struct {
	Id             uint8     `json:"id"`
	Title          string    `json:"title"`
	Image          string    `json:"image"`
	Description    string    `json:"descripton"`
	Video_url      string    `json:"video_url"`
	EquipmentId    Equipment `json:"equipment"`
	Sets           uint8     `json:"sets"`
	Reps           uint8     `json:"reps"`
	Difficult      uint8     `json:"difficult"`
	Lead_muscle_id Muscles   `json:"muscle"`
}
