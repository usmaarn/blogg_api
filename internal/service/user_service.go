package service

import (
	"fmt"
	"strconv"

	"github.com/usmaarn/blogg_api/cmd/config"
	"github.com/usmaarn/blogg_api/internal/dto"
	model "github.com/usmaarn/blogg_api/internal/models"
	"github.com/usmaarn/blogg_api/package/enum"
	"github.com/usmaarn/blogg_api/package/utils"
)

func CreateUser(dto dto.CreateUserDto) (model.User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Name:         fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
		EmailAddress: dto.Email,
		PhoneNumber:  dto.PhoneNumber,
		Type:         strconv.Itoa(enum.UserTypeBasic),
		Status:       strconv.Itoa(enum.UserStatusActive),
		Password:     hashedPassword,
	}

	result := config.DB.Create(&user)
	return user, result.Error
}

func FindUserById(userId uint) (model.User, error) {
	var user model.User
	result := config.DB.First(&user)
	return user, result.Error
}

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
