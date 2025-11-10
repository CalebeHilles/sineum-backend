package routes

import (
	controller "api-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/blogs", controller.GetEveryBlogs).Methods("GET")
	r.HandleFunc("/api/blogs/{id}", controller.GetBlogsById).Methods("GET")
	r.HandleFunc("/api/blogs", controller.CreateBlogs).Methods("POST")
	r.HandleFunc("/api/blogs/{id}", controller.DeleteBlogs).Methods("DELETE")
	r.HandleFunc("/api/blogs/{id}", controller.EditBlogs).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
