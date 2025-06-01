package registerhandler

import (
	"strconv"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func DeleteQuoteByID(router *router.Router) {

	router.HandleFunc(
		"DELETE /quotes/{id}",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			id := ctx.Request().PathValue("id")
			num, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}
			err = ctx.Services().Quotes().DeleteQuote(service.DeleteQuoteInput{
				ID: num,
			})
			if err != nil {
				return nil, err
			}
			return "successed", nil
		},
	)
}
