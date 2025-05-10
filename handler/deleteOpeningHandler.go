package handler

import (
	"fmt"
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	openingId := ctx.Query("id")
	if err := validateHasOpeningId(openingId); err != nil {
		dto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	openingSchema, findOpeningError := findOpeningById(ctx, openingId)
	if findOpeningError != nil {
		return
	}
	if err := deleteOpening(openingSchema); err != nil {
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendSuccessResponse(ctx, "delete-opening", http.StatusOK, openingSchema)
}

func deleteOpening(opening *schemas.Opening) error {
	if err := db.Delete(&opening).Error; err != nil {
		return fmt.Errorf("Error: failed to delete opening with id: %s", opening.ID)
	}
	return nil
}
