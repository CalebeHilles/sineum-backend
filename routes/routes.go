package routes

import (
	controller "api-go/controllers"
	"api-go/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/blogs", controller.GetEveryBlogs).Methods("GET")
	r.HandleFunc("/api/blogs/{id}", controller.GetBlogsById).Methods("GET")
	r.HandleFunc("/api/blogs", controller.CreateBlogs).Methods("POST")
	r.HandleFunc("/api/blogs/{id}", controller.DeleteBlogs).Methods("DELETE")
	r.HandleFunc("/api/blogs/{id}", controller.EditBlogs).Methods("PUT")

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
			"https://kleb-letter.vercel.app",
		}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "PUT", "DELETE",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type", "Authorization",
		}),
	)

	log.Fatal(http.ListenAndServe(":8000", corsOptions(r)))
}
