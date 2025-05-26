package service

import (
	"github.com/usmaarn/blogg_api/internal/dto"
	model "github.com/usmaarn/blogg_api/internal/models"
)

func RegisterUser(dto dto.CreateUserDto) (model.Session, error) {
	user, err := CreateUser(dto)
	if err != nil {
		return model.Session{}, err
	}

	return CreateSession(user.ID)
}
