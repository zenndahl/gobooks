package handlers

import (
	"gobooks/schemas"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowRandomBookHandler(ctx *gin.Context) {
	books := []schemas.Book{}
	var bookCount int
	count := int64(bookCount)

	db.Model(&books).Count(&count)
	bookCount = int(count)

	logger.Info(bookCount)
	num := strconv.Itoa(rand.Intn(bookCount))

	id := ctx.Query(num)
	book := schemas.Book{}
	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error finding book")
		return
	}

	sendSuccess(ctx, "show-book", book)
}
