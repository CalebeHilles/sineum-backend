package routes

import (
	controller "api-go/controllers"
	"api-go/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/blogs", controller.GetEveryBlogs).Methods("GET")
	r.HandleFunc("/blogs/{id}", controller.GetBlogsById).Methods("GET")
	r.HandleFunc("/blogs", controller.CreateBlogs).Methods("POST")
	r.HandleFunc("/blogs/{id}", controller.DeleteBlogs).Methods("DELETE")
	r.HandleFunc("/blogs/{id}", controller.EditBlogs).Methods("PUT")

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
			"https://sineum.vercel.app",
		}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "PUT", "DELETE",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type", "Authorization",
		}),
	)

	port := os.Getenv("PORT")
	println("Project running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, corsOptions(r)))
}
