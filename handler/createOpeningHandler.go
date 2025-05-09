package handler

import (
	"fmt"
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request, err := parseAndValidateRequest(ctx)
	if err != nil {
		dto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := mapToOpeningModel(request)

	if err := db.Create(&opening).Error; err != nil {
		logger.Error("Failed to create opening: %v", err.Error())
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, constant.ErrCreateOpenings+err.Error())
		return
	}

	dto.SendSuccessResponse(ctx, "create-opening", http.StatusCreated, opening)
}

func parseAndValidateRequest(ctx *gin.Context) (*dto.CreateOpeningRequest, error) {
	request := &dto.CreateOpeningRequest{}
	if err := ctx.BindJSON(request); err != nil {
		logger.Error("Failed to bind JSON: %v", err.Error())
		return nil, fmt.Errorf(constant.ErrInvalidJson)
	}

	if err := request.Validate(); err != nil {
		logger.ErrorFormated("Validation error when creating opening: %v", err.Error())
		return nil, fmt.Errorf(constant.ErrValidateOpenings + err.Error())
	}

	return request, nil
}

func mapToOpeningModel(request *dto.CreateOpeningRequest) schemas.Opening {
	return schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
}
