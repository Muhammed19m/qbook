package router

import (
	"net/http"

	"github.com/Muhammed19m/qbook/internal/controller/http2"
)

type Router struct {
	Service http2.Services
	http.ServeMux
}

func (r *Router) HandleFunc(pattern string, middlewares []http2.Middleware, handler http2.HandleFunc) {
	handlerRW := http2.WrapHandlerInMiddlewares(handler, middlewares...)
	newHandler := r.modulation(handlerRW)
	r.ServeMux.HandleFunc(pattern, newHandler)

}
