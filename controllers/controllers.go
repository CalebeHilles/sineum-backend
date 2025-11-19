package controller

import (
	"api-go/database"
	"api-go/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEveryBlogs(w http.ResponseWriter, r *http.Request) {
	var b []models.Blog
	database.DB.Find(&b)
	json.NewEncoder(w).Encode(b)
}

func GetBlogsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var b models.Blog
	database.DB.First(&b, id)
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
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var b []models.Blog
	database.DB.Delete(&b, id)
	json.NewEncoder(w).Encode(b)
}

func EditBlogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var b models.Blog
	database.DB.First(&b, id)
	json.NewDecoder(r.Body).Decode(&b)
	database.DB.Save(&b)
	json.NewEncoder(w).Encode(b)
}
