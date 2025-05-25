package handler

import (
	"net/http"

	"github.com/usmaarn/blogg_api/internal/dto"
	"github.com/usmaarn/blogg_api/internal/service"
	"github.com/usmaarn/blogg_api/package/utils"
	"github.com/usmaarn/blogg_api/package/utils/response"
	"github.com/usmaarn/blogg_api/package/validators"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestDto dto.CreateUserDto
	err := utils.BindJSON(r.Body, &requestDto)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	errorMap, err := validators.ValidateStruct(requestDto)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "", errorMap)
		return
	}

	user, err := service.CreateUser(requestDto)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.SuccessResponse(w, http.StatusCreated, "", user)
}
