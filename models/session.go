package models

import "time"

type Session struct {
	ID          int64     `json:"-"`
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"-"`
	IpAddress   *string   `json:"-"`
	UserAgent   *string   `json:"-"`
}
