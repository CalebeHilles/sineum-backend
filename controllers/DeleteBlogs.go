package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var b models.Blog
	fetchResult := database.DB.First(&b, id)

	if fetchResult.Error != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	result := database.DB.Delete(&b, id)

	if result.Error != nil {
		http.Error(w, "Server failed deleting this blog", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(b)
}
