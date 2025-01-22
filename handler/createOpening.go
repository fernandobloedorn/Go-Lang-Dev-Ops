package handler

import (
	"net/http"

	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// logger.Infof("Request received: %+v", request)

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)

}
