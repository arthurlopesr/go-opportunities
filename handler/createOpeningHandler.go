package handler

import (
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := dto.CreateOpeningRequest{}
	
	if err := ctx.BindJSON(&request); err != nil {
		logger.Error("Failed to bind JSON: %v", err.Error())
		dto.SendErrorResponse(ctx, http.StatusBadRequest, constant.ErrInvalidJson)
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrorFormated("validation error when creating opening", err.Error())
		dto.SendErrorResponse(ctx, http.StatusBadRequest, constant.ErrValidateOpenings+err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Error("Failed to create opening: %v", err.Error())
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, constant.ErrCreateOpenings+err.Error())
		return
	}
	dto.SendSuccessResponse(ctx, "create-opening", http.StatusCreated, opening)
}
