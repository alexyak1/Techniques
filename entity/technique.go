package entity

type Technique struct {
	Id string `json:"id"`
    Name string `json:"name"`
    Belt string `json:"belt"`
    ImageURL string `json:"image_url"`
    Type string `json:"type"`
}