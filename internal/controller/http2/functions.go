package http2

import (
	"net/http"
	"strconv"
)

func AdaptToRW(handle HandleFunc) HandleFuncRW {
	return func(ctx ContextRW) (any, error) {
		return handle(ctx)
	}
}

func WrapHandlerInMiddlewares(handler HandleFunc, middlewares ...Middleware) HandleFuncRW {
	hmw := AdaptToRW(handler)

	for _, mw := range middlewares {
		hmw = mw(hmw)
	}
	return hmw
}

func PathInt(r *http.Request, name string) (int, error) {
	return strconv.Atoi(r.PathValue(name))
}
