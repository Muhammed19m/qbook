package http2

import (
	"net/http"

	"github.com/Muhammed19m/qbook/internal/service"
)

type Router interface {
	HandleFunc(pattern string, chainMW []Middleware, handler HandleFunc)
}

type HandleFunc func(Context) (any, error)

type HandleFuncRW func(ContextRW) (any, error)

type Context interface {
	RequestID() string

	Request() *http.Request

	Writer() http.ResponseWriter

	Services() Services
}

type ContextRW interface {
	Context
	SetRequestID(string)
	SetRequest(*http.Request)
}

type Services interface {
	Quotes() *service.Quotes
}

type Middleware func(HandleFuncRW) HandleFuncRW
