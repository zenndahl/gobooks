package handlers

import (
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List book
// @Description List all books
// @Tags Books
// @Accept json
// @Produce json
// @Param request body BookRequest true "Request body"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func ListBooksHandler(ctx *gin.Context) {
	book := []schemas.Book{}
	if err := db.Find(&book).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing books")
		return
	}
	sendSuccess(ctx, "list-books", book)
}
