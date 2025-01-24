package handler

import (
	"fmt"
	"net/http"

	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /apu/v1

// @Summary Delete opening
// @Description Delete a opening job
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deletiong with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
