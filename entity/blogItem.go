package entity

type BlogItem struct {
	ID      int     `json:"Id"`
	Title   string  `json:"Title"`
	Date    string  `json:"date"`
	Option1 Option1 `json:"option1"`
	Option2 Option2 `json:"option2"`
	Option3 Option3 `json:"option3"`
	Option4 Option4 `json:"option4"`
}
type Option1 struct {
	Name  string `json:"name"`
	Votes string `json:"votes"`
}
type Option2 struct {
	Name  string `json:"name"`
	Votes string `json:"votes"`
}
type Option3 struct {
	Name  string `json:"name"`
	Votes string `json:"votes"`
}
type Option4 struct {
	Name  string `json:"name"`
	Votes string `json:"votes"`
}
