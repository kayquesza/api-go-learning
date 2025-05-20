package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/api-go-learning/schemas"
)

// @BasePath /api/v1

// @Summary      List Opening
// @Description  List a job opening
// @Tags         opening
// @Accept       json
// @Produce      json
// @Success      200      {object}  ListOpeningsResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /opening [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
