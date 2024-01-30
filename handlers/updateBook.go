
package handlers

import (
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookHandler(ctx *gin.Context) {
	request := UpdateBookRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "book not found")
	}

	if request.Name != "" {
		book.Name = request.Name
	}
	if request.Author != "" {
		book.Author = request.Author
	}
  if request.Genre != "" {
		book.Genre = request.Genre
	}
	if request.Year != "" {
		book.Year = request.Year
	}

  if err := db.Save(&book).Error; err != nil {
		logger.Errf("error updating book: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating book")
		return
	}

	sendSuccess(ctx, "update-book", book)
}
