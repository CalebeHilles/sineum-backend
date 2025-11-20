package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func CreateBlogs(w http.ResponseWriter, r *http.Request) {
	var newBlog models.Blog
	json.NewDecoder(r.Body).Decode(&newBlog)
	database.DB.Create(&newBlog)
	json.NewEncoder(w).Encode(newBlog)
}

func DeleteBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var b []models.Blog
	database.DB.Delete(&b, id)
	json.NewEncoder(w).Encode(b)
}

func EditBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var b models.Blog
	database.DB.First(&b, id)
	json.NewDecoder(r.Body).Decode(&b)
	database.DB.Save(&b)
	json.NewEncoder(w).Encode(b)
}
