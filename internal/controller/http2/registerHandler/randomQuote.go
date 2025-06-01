package registerhandler

import (
	"fmt"
	"log/slog"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/domain"
)

func RandomQuote(router *router.Router) {

	router.HandleFunc(
		"GET /quotes/random",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			quote, err := ctx.Services().Quotes().GetRandomQuote()
			if err != nil {
				return nil, err
			}
			// todo:
			slog.Info(fmt.Sprint(quote))
			var zeroQuote domain.Quote
			if quote == zeroQuote {
				return nil, nil
			}
			return quote, nil
		},
	)
}
