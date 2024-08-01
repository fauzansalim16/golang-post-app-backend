package routes

import (
	"backend-go/handlers"
	"backend-go/repositories"
	"backend-go/services"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	postRepo := &repositories.PostRepository{}
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	r.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/posts", postHandler.ShowPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", postHandler.DeletePost).Methods("DELETE")

	fs := http.FileServer(http.Dir("./uploads"))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	return r
}
