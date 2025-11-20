package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"
)

func CreateBlogs(w http.ResponseWriter, r *http.Request) {
	var newBlog models.Blog
	json.NewDecoder(r.Body).Decode(&newBlog)
	result := database.DB.Create(&newBlog)

	if result.Error != nil {
		http.Error(w, "Server failed creating blog", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newBlog)
}
