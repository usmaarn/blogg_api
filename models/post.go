package models

import (
	"html/template"
	"time"
)

type Post struct {
	ID        int64
	Title     string
	UserID    int64
	Content   template.HTML
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePost struct {
	Title   string        `validate:"required,min=3"`
	Content template.HTML `validate:"required,min=10"`
}
