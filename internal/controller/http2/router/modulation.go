package router

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
)

func (router *Router) modulation(handlerRW http2.HandleFuncRW) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			respData any
			data     []byte
			err      error
		)

		if err = r.ParseForm(); err != nil {
			slog.Warn("modulation.Request.ParseForm: ", err.Error())
		}

		respData, err = handlerRW(&rwContext{
			request:  r,
			writer:   w,
			services: router.Service,
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respData = ResponseError{
				Error: err.Error(),
			}
		}

		if s, ok := respData.(string); ok {
			respData = ResponseMsg{
				Message: s,
			}
		}

		if data, err = json.Marshal(respData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Error("modulation.json.Marshal: " + err.Error())
			data = fmt.Appendf(nil, `{"error":"%v"}`, err)
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Error("modulation.w.Write: " + err.Error())
		}
	}
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseMsg struct {
	Message string `json:"message"`
}
