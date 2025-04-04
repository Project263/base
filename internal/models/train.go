package models

type Train struct {
	Id string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Image     string `json:"image"`
	Video_url string `json:"video_url"`
}

type TrainWithMuscle struct {
	Id string `json:"id"`

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Image     string `json:"image,omitempty"`
	Video_url string `json:"video_url,omitempty"`

	Muscles Muscle `json:"muscles"`
}
