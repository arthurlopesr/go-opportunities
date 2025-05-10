package handler

import (
	"github.com/arthurlopesr/go-opportunities/dto"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllOpeningHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		dto.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendSuccessResponse(ctx, "find-all-openings", http.StatusOK, openings)
}
