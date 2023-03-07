package models

type Album struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Content string `json:"content,omitempty"`
}
