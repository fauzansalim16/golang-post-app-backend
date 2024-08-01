package services

import (
	"backend-go/models"
	"backend-go/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	// Tambahkan validasi atau logika bisnis lainnya di sini
	return s.repo.CreatePost(post)
}

func (s *PostService) ShowPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
