package registerhandler

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/domain"
)

func RandomQuote(router *router.Router) {
	var zeroQuote domain.Quote
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
			// ----------------------
			var zeroQuote domain.Quote
			if quote == zeroQuote {
				return nil, nil
			}
			// ----------------------
			if quote == (domain.Quote{}) {
				return nil, nil
			}
			// ----------------------
			if quote.IsZero() {
				return "нет цитат", nil
			}
			// ----------------------
			if isZeroQuote(quote) {
				return nil, errors.New("нет цитат")
			}

			return quote, nil
		},
	)
}

func isZeroQuote(q domain.Quote) bool {
	return q == (domain.Quote{})
}
