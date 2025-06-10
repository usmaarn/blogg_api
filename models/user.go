package models

import "time"

type User struct {
	ID        int64 `db:"name"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
