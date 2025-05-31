package router

import (
	"net/http"

	"github.com/Muhammed19m/qbook/internal/service"
)

type Router struct {
	Service *service.Quotes
	http.ServeMux
}

func (c *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	c.ServeMux.HandleFunc(pattern, handlerFunc)
}

