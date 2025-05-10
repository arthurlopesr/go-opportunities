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

func updateOpening(opening *schemas.Opening, request *dto.UpdateOpeningRequest, ctx *gin.Context) (*schemas.Opening, error) {
	updateField := func(field interface{}, value interface{}) {
		switch v := value.(type) {
		case string:
			if v != "" {
				*field.(*string) = v
			}
		case *bool:
			if v != nil {
				*field.(*bool) = *v
			}
		case int64:
			if v > 0 {
				*field.(*int64) = v
			}
		}
	}

	updateField(&opening.Role, request.Role)
	updateField(&opening.Company, request.Company)
	updateField(&opening.Location, request.Location)
	updateField(&opening.Remote, request.Remote)
	updateField(&opening.Link, request.Link)
	updateField(&opening.Salary, request.Salary)

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorFormated("error updating opening: %v", err.Error())
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, "error updating opening")
		return nil, err
	}
	return opening, nil
}
