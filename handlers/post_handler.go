package handlers

import (
	"backend-go/models"
	"backend-go/services"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get form values
	title := r.FormValue("title")
	description := r.FormValue("description")

	// Get the file from form data
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a unique filename
	filename := uuid.New().String() + filepath.Ext(header.Filename)

	// Define the path where the file will be saved
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filePath := filepath.Join(uploadDir, filename)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the uploaded file to the created file on the filesystem
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create post with image path
	post := &models.Post{
		Title:       title,
		Description: description,
		Image_url:   filePath,
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) ShowPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.ShowPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	// Ekstrak ID dari URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Konversi ID dari string ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Panggil service untuk menghapus post
	err = h.postService.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post deleted successfully"))
}
