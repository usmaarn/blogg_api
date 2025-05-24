package entites

import (
	"time"
)

var Users = []User{}

type User struct {
	ID              int
	Name            string
	EmailAddress    string
	PhoneNumber     string
	Password        string
	Role            string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (user User) String() string {
	return user.EmailAddress
}

func (user *User) HasVerifiedEmail() bool {
	return !user.EmailVerifiedAt.IsZero()
}

func (user *User) Save() bool {
	Users = append(Users, *user)
	return true
}

func (user *User) Delete() bool {
	users := []User{}
	for _, el := range Users {
		if el.ID != user.ID {
			users = append(users, el)
		}
	}
	Users = users
	return true
}
