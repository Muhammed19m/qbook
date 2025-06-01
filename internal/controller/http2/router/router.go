package router

import (
	"net/http"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
)

type Router struct {
	Service http2.Services
	http.ServeMux
}

func (r *Router) HandleFunc(pattern string, midllewares []http2.Middleware, handler http2.HandleFunc) {
	handlerRW := http2.WrapHandlerInMidllewares(handler, midllewares...)
	newHandler := r.modulation(handlerRW)
	r.ServeMux.HandleFunc(pattern, newHandler)

}
