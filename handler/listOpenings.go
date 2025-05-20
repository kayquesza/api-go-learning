package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/api-go-learning/schemas"
)

// @BasePath /api/v1

// @Summary      List Openings
// @Description  List all job openings
// @Tags         Openings
// @Accept       json
// @Produce      json
// @Success      200      {object}  ListOpeningsResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
