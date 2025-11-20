package models

type Blog struct {
	Id          string    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var Blogs []Blog
