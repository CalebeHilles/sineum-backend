package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"
)

func GetEveryBlogs(w http.ResponseWriter, r *http.Request) {
	var b []models.Blog
	result := database.DB.Find(&b)

	if result.Error != nil {
		http.Error(w, "Server failed getting every blogs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(b)
}
