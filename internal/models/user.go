package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" db:"id"`
	Name            string         `json:"name" db:"name"`
	Username        *string        `json:"username" db:"username"`
	EmailAddress    string         `json:"emailAddress" db:"email" gorm:"column:email"`
	PhoneNumber     string         `json:"phoneNumber" db:"phone_number"`
	Password        string         `json:"-" db:"password"`
	Type            string         `json:"type" db:"type"`
	Status          string         `json:"status" db:"status"`
	EmailVerifiedAt *time.Time     `json:"emailVerifiedAt" db:"email_verified_at"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" db:"deleted_at"`
}

func (user User) String() string {
	return user.EmailAddress
}

func (user *User) HasVerifiedEmail() bool {
	return !user.EmailVerifiedAt.IsZero()
}
