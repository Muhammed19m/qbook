package registerhandler

import (
	"encoding/json"
	"net/http"

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
		func(w http.ResponseWriter, r *http.Request) {
			var rb requestBody
			if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
				w.Write([]byte("error decode body"))
				return
			}

			in := service.AddQuoteInput{
				Author: rb.Author,
				Text:   rb.Quote,
			}
			quote, err := router.Service.AddQuote(in)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(quote)
		},
	)
}
