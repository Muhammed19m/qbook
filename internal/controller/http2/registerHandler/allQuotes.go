package registerhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func AllQuotes(router *router.Router) {

	router.HandleFunc(
		"GET /quotes",
		func(w http.ResponseWriter, r *http.Request) {

			author := r.URL.Query().Get("author")
			if author == "" {

				quote, err := router.Service.AllQuotes()
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(quote)
			} else {
				in := service.QuoteByAuthorInput{
					Author: author,
				}

				quotes, err := router.Service.QuoteByAuthor(in)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(quotes)
			}
		},
	)
}
