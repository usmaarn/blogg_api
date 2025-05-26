package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmaarn/blogg_api/internal/dto"
	"github.com/usmaarn/blogg_api/internal/service"
	"github.com/usmaarn/blogg_api/package/utils"
	"github.com/usmaarn/blogg_api/package/utils/response"
)

func RegisterHandler(c *gin.Context) {
	var requestDto dto.CreateUserDto
	if err := c.ShouldBindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.ValidateJson(err, requestDto))
		return
	}

	session, err := service.RegisterUser(requestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse("", session))
}

func LoginHandler(c *gin.Context) {
	var dto dto.LoginDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.ValidateJson(err, dto))
		return
	}

	user, err := service.FindUserByEmail(dto.EmailAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("invalid email or password", nil))
		return
	}

	if isCorrectPassword := utils.VerifyPassword(dto.Password, user.Password); !isCorrectPassword {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("invalid email or password", nil))
		return
	}

	session, err := service.CreateSession(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("invalid email or password", nil))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("created", session))
}
