package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/api-go-learning/schemas"
)

// @BasePath /api/v1

// @Summary      Show Opening
// @Description  Show a job opening
// @Tags         Openings
// @Accept       json
// @Produce      json
// @Param id query string true "Opening Identification"
// @Success      200      {object}  ShowOpeningResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParanIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
