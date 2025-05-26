package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/usmaarn/blogg_api/cmd/config"
	model "github.com/usmaarn/blogg_api/internal/models"
)

func CreateSession(userID uint) (model.Session, error) {
	session := model.Session{
		UserID:       userID,
		AccessToken:  uuid.NewString(),
		RefreshToken: uuid.NewString(),
		ExpiresAt:    time.Now().Add(30 * 24 * time.Hour),
	}

	result := config.DB.Create(&session)
	return session, result.Error
}

func ValidateAccessToken(token string, userID uint) bool {
	var session model.Session
	result := config.DB.Where("access_token = ?", token).First(&session)
	if result.Error != nil {
		return false
	}

	if session.ExpiresAt.Before(time.Now()) {
		config.DB.Delete(&session)
		return false
	}

	return session.UserID == userID
}
