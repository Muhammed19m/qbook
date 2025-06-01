package registerhandler

import (
	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func GetQuotes(router *router.Router) {
	router.HandleFunc(
		"GET /quotes",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			in := service.QuotesInput{
				Author: ctx.Request().URL.Query().Get("author"),
			}

			return ctx.Services().Quotes().Quotes(in)
		},
	)
}
