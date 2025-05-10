package handler

import (
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindOpeningByIdHandler(ctx *gin.Context) {
	openingId := ctx.Query("id")
	if err := validateHasOpeningId(openingId); err != nil {
		dto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening, err := findOpeningById(ctx, openingId)
	if err != nil {
		return
	}
	dto.SendSuccessResponse(ctx, "find-opening", http.StatusOK, opening)
}

func findOpeningById(ctx *gin.Context, openingId string) (*schemas.Opening, error) {
	opening := schemas.Opening{}
	if err := db.First(&opening, openingId).Error; err != nil {
		dto.SendErrorResponse(ctx, http.StatusNotFound, "error: opening not found")
		return nil, err
	}
	return &opening, nil
}
