package handlers

import (
	"gobooks/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(ctx *gin.Context) {
	request := CreateBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %+v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

  book := schemas.Book{
    Name:   request.Name,
    Author: request.Author,
    Genre:  request.Genre,
    Year:   request.Year,
	}

	if err := db.Create(&book).Error; err != nil {
		logger.Errf("error creating book: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating book on database")
		return
	}

	sendSuccess(ctx, "create-book", book)
}
