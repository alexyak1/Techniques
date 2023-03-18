package entity

type Technique struct {
	Id       *string `json:"id"`
	Name     string  `json:"name"`
	Belt     string  `json:"belt"`
	ImageURL string  `json:"image_url"`
	ImageId  string  `json:"image_id"`
	Type     string  `json:"type"`
}

type KataTechnique struct {
	Id       *string `json:"id"`
	Name     string  `json:"name"`
	KataName string  `json:"kata_name"`
	ImageURL string  `json:"image_url"`
	Type     string  `json:"type"`
	ImageId  string  `json:"image_id"`
}
