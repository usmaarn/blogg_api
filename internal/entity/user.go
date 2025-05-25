package entity

import (
	"time"
)

type User struct {
	ID              int64     `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	EmailAddress    string    `json:"emailAddress" db:"email"`
	PhoneNumber     string    `json:"phoneNumber" db:"phone_number"`
	Password        string    `json:"-" db:"password"`
	Type            int       `json:"type" db:"type"`
	Status          int       `json:"status" db:"status"`
	EmailVerifiedAt time.Time `json:"emailVerifiedAt" db:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

func (user User) String() string {
	return user.EmailAddress
}

func (user *User) HasVerifiedEmail() bool {
	return !user.EmailVerifiedAt.IsZero()
}

func (user *User) Save() User {
	return *user
}

func (user *User) Delete() bool {
	return true
}
