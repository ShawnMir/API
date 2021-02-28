package handler

// Articles ...
var Articles []Article

// Article ...
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
