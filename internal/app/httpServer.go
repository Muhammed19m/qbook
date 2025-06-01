package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	registerhandler "github.com/Muhammed19m/qbook/internal/controller/http2/registerHandler"
	"github.com/Muhammed19m/qbook/internal/controller/http2/router"

	"golang.org/x/sync/errgroup"
)

func initHttpServer(ss *services) *http.Server {
	r := &router.Router{
		Service: ss,
	}
	registerHandlers(r)

	addr := flag.String("addr", ":8080", "addr to listen on")
	flag.Parse()

	return &http.Server{
		Addr:    *addr,
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
