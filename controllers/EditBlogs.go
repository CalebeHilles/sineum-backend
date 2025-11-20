package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func EditBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var b models.Blog
	fetchResult := database.DB.First(&b, id)

	if fetchResult.Error != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&b)
	result := database.DB.Save(&b)

	if result.Error != nil {
		http.Error(w, "Server failed saving changes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(b)
}
