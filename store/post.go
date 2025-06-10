package store

import (
	"database/sql"

	"github.com/usmaarn/blogg/models"
)

func (s *Store) SavePost(post *models.Post) error {
	var row *sql.Row
	if post.ID == 0 {
		query := `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
		row = s.DB.QueryRow(query, post.UserID, post.Title, post.Content)
	} else {
		query := `UPDATE posts SET title = $1, content = $2, updated_at = NOW() WHERE id = $3 RETURNING id, created_at, updated_at`
		row = s.DB.QueryRow(query, post.Title, post.Content, post.ID)
	}

	if row.Err() != nil {
		return row.Err()
	}

	return row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

func (s *Store) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := s.DB.Query("SELECT id, user_id, title, content, created_at, updated_at FROM posts")
	if err != nil {
		return []models.Post{}, err
	}

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
