package registerhandler

import (
	"encoding/json"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func AddQuote(router *router.Router) {

	type requestBody struct {
		Author string `json:"author"`
		Quote  string `jsom:"quote"`
	}

	router.HandleFunc(
		"POST /quotes",
		[]http2.Middleware{},
		func(ctx http2.Context) (any, error) {
			var rb requestBody
			if err := json.NewDecoder(ctx.Request().Body).Decode(&rb); err != nil {
				return nil, err
			}

			in := service.AddQuoteInput{
				Author: rb.Author,
				Text:   rb.Quote,
			}
			quote, err := ctx.Services().Quotes().AddQuote(in)
			if err != nil {
				return nil, err
			}

			return quote, nil
		},
	)
}
