package repositories

import (
	"backend-go/config"
	"backend-go/models"
	"log"
	"os"
)

type PostRepository struct{}

func (r *PostRepository) CreatePost(post *models.Post) error {
	query := "INSERT INTO posts (title, description, image_url) VALUES (?, ?, ?)"
	_, err := config.DB.Exec(query, post.Title, post.Description, post.Image_url)
	return err
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	rows, err := config.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Image_url, &post.Posted_at); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) DeletePost(id int) error {
	// Pertama, dapatkan path gambar (jika ada)
	var imagePath string
	err := config.DB.QueryRow("SELECT image_url FROM posts WHERE id = ?", id).Scan(&imagePath)
	if err != nil {
		return err
	}

	// Hapus record dari database
	_, err = config.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}

	// Jika ada file gambar, hapus juga
	if imagePath != "" {
		err = os.Remove(imagePath)
		if err != nil {
			// Log error tapi jangan hentikan proses
			log.Printf("Error deleting image file: %v", err)
		}
	}

	return nil
}
