package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBlogsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var b models.Blog
	result := database.DB.First(&b, id)

	if result.Error != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(b)
}
