package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	registerhandler "github.com/Muhammed19m/qbook/internal/controller/http2/registerHandler"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"

	"golang.org/x/sync/errgroup"
)

type HttpServerConfig struct {
	Addr string
}

func initHttpServer(ss *services, cfg HttpServerConfig) *http.Server {
	r := &router.Router{
		Service: ss,
	}
	registerHandlers(r)

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}
}

func runHttpServer(ctx context.Context, server *http.Server) error {
	g, ctx := errgroup.WithContext(ctx)

	// Запуск сервера
	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server.ListenAndServe: %w", err)
		}
		return nil
	})

	// Завершение сервера при завершении контекста
	g.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(ctx)
	})

	return g.Wait()
}

func registerHandlers(r *router.Router) {
	// /quotes

	registerhandler.AddQuote(r)        // 1. Добавление новой цитаты
	registerhandler.GetQuotes(r)       // 2. Получение цитат
	registerhandler.RandomQuote(r)     // 3. Получение случайной цитаты
	registerhandler.DeleteQuoteByID(r) // 5. Удаление цитаты
}
