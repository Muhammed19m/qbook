package registerhandler

import (
	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func AllQuotes(router *router.Router) {

	router.HandleFunc(
		"GET /quotes",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			author := ctx.Request().URL.Query().Get("author")
			if author == "" {
				quotes, err := ctx.Services().Quotes().AllQuotes()
				if err != nil {
					return nil, err
				}
				return quotes, nil
			} else {
				in := service.QuoteByAuthorInput{
					Author: author,
				}

				quotes, err := ctx.Services().Quotes().QuoteByAuthor(in)
				if err != nil {
					return nil, err
				}
				return quotes, nil
			}
		},
	)
}
