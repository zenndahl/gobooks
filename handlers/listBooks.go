
package handlers

import (
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooksHandler(ctx *gin.Context) {
	book := []schemas.Book{}
	if err := db.Find(&book).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing books")
		return
	}
	sendSuccess(ctx, "list-books", book)
}
