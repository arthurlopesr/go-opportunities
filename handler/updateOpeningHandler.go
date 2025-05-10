package handler

import (
	"fmt"
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request, err := parseAndValidateUpdateRequest(ctx)
	if err != nil {
		dto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openingId := ctx.Query("id")
	if err := validateHasOpeningId(openingId); err != nil {
		dto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openingSchema, findOpeningError := findOpeningById(ctx, openingId)
	if findOpeningError != nil {
		return
	}

	if _, err := updateOpening(openingSchema, request, ctx); err != nil {
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendSuccessResponse(ctx, "update-opening", http.StatusOK, openingSchema)
}

func parseAndValidateUpdateRequest(ctx *gin.Context) (*dto.UpdateOpeningRequest, error) {
	request := dto.UpdateOpeningRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		logger.Error("Failed to bind JSON: %v", err.Error())
		return nil, fmt.Errorf(constant.ErrInvalidJson)
	}
	if err := request.Validate(); err != nil {
		logger.ErrorFormated("Validation error when updating opening: %v", err.Error())
		return nil, fmt.Errorf(constant.ErrValidateOpenings + err.Error())
	}
	return &request, nil
}

func updateOpening(opening *schemas.Opening, request *dto.UpdateOpeningRequest, ctx *gin.Context) (openingUpdated *schemas.Opening, err error) {
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorFormated("error updating opening: %v", err.Error())
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, "error updating opening")
		return nil, err
	}
	return opening, nil
}
