package handler

import (
	"net/http"

	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /apu/v1

// @Summary List openings
// @Description List openings job
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listinig openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
