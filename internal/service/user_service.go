package service

import (
	"fmt"

	"github.com/usmaarn/blogg_api/cmd/config"
	"github.com/usmaarn/blogg_api/internal/dto"
	entity "github.com/usmaarn/blogg_api/internal/entity"
	"github.com/usmaarn/blogg_api/package/enum"
	"github.com/usmaarn/blogg_api/package/utils"
)

func CreateUser(dto dto.CreateUserDto) (entity.User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return entity.User{}, err
	}

	var user = entity.User{
		Name:         fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
		EmailAddress: dto.Email,
		PhoneNumber:  dto.PhoneNumber,
		Type:         enum.UserTypeBasic,
		Status:       enum.UserStatusActive,
		Password:     hashedPassword,
	}

	row := config.DB.QueryRow(
		`INSERT INTO users (name, email, phone_number, type, status, password)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id;`,
		user.Name, user.EmailAddress, user.PhoneNumber, user.Type, user.Status, user.Password,
	)

	var userId int64
	err = row.Scan(&userId)
	if err != nil {
		return entity.User{}, err
	}
	return GetUserById(userId)
}

func GetUserById(userId int64) (entity.User, error) {
	var user entity.User
	err := config.DB.Select(&user, "SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
