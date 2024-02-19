package handlers

import (
	"fmt"
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete book
// @Description Delete a book entry
// @Tags Books
// @Accept json
// @Produce json
// @Param request body BookRequest true "Request body"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /book [delete]
func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}
	book := schemas.Book{}
	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id %s not found", id))
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting book with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-book", book)
}
