package models

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Contnet string `json:"content"`
	Author  string `json:"author,omitempty"`
}
