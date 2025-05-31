package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	registerhandler "github.com/Muhammed19m/qbook/internal/controller/http2/registerHandler"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"

	"github.com/Muhammed19m/qbook/internal/service"
	"golang.org/x/sync/errgroup"
)

func initHttpServer(s *service.Quotes) *http.Server {
	r := &router.Router{
		Service: s,
	}
	registerHandlers(r)

	return &http.Server{
		Addr:    ":8080",
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

	// 1. Добавление новой цитаты (POST /quotes)
	registerhandler.AddQuote(r)

	// 2. Получение всех цитат (GET /quotes)
	// 4. Фильтрация по автору (GET /quotes?author=Confucius)
	registerhandler.AllQuotes(r)

	// 3. Получение случайной цитаты (GET /quotes/random)
	registerhandler.RandomQuote(r)

	// 5. Удаление цитаты по ID (DELETE /quotes/{id})
	registerhandler.DeleteQuoteByID(r)
}
