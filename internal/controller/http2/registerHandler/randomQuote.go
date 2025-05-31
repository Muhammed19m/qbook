package registerhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
)

func RandomQuote(router *router.Router) {

	router.HandleFunc(
		"GET /quotes/random",
		func(w http.ResponseWriter, r *http.Request) {

			quote, err := router.Service.GetRandomQuote()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(quote)
		},
	)
}
