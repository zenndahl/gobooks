package handlers

import (
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show book
// @Description Show a specific book entry
// @Tags Books
// @Accept json
// @Produce json
// @Param request body BookRequest true "Request body"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book [get]
func ShowBookHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}
	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error finding book")
		return
	}

	sendSuccess(ctx, "show-book", book)
}
