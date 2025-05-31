package registerhandler

import (
	"net/http"
	"strconv"

	"github.com/Muhammed19m/qbook/internal/controller/http2/router"
	"github.com/Muhammed19m/qbook/internal/service"
)

func DeleteQuoteByID(router *router.Router) {

	router.HandleFunc(
		"DELETE /quotes/{id}",
		func(w http.ResponseWriter, r *http.Request) {

			id := r.PathValue("id")
			num, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "invalid parametr {id}", http.StatusBadRequest)
				return
			}
			err = router.Service.DeleteQuote(service.DeleteQuoteInput{
				ID: num,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			// w.Write([]byte("success"))

		},
	)
}
