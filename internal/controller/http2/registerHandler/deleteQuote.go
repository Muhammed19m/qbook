package registerhandler

import (
	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func DeleteQuoteByID(router *router.Router) {

	router.HandleFunc(
		"DELETE /quotes/{id}",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			num, _ := http2.PathInt(ctx.Request(), "id")

			// Удалить цитату
			in := service.DeleteQuoteInput{
				ID: num,
			}

			return nil, ctx.Services().Quotes().DeleteQuote(in)
		},
	)
}
