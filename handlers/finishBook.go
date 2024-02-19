package handlers

import (
	"gobooks/schemas"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FinishBookHandler(ctx *gin.Context) {
	request := FinishBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %+v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	finishedBook := schemas.FinishedBook{
		Name:       request.Name,
		Author:     request.Author,
		Genre:      request.Genre,
		Year:       request.Year,
		Finished:   request.Finished,
		FinishedAt: request.FinishedAt,
		Rating:     request.Rating,
	}

	finishedBook.FinishedAt = time.Now()

	if err := db.Create(&finishedBook).Error; err != nil {
		logger.Errf("error marking book as fnished: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating book on database")
		return
	}

	sendSuccess(ctx, "finished-book", finishedBook)
}
