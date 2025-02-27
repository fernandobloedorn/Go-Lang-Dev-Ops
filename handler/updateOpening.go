package handler

import (
	"net/http"

	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /apu/v1

// @Summary Update opening
// @Description Update a opening job
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Param opening body UpdateOpeningRequest true "Opeing data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.Find(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not foud")
		return
	}

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
		logger.Errorf("Erros updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, "success updating openig", opening)
}
