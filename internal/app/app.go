package app

import (
	"context"

	"github.com/Muhammed19m/qbook/internal/repository/memory"
	"github.com/Muhammed19m/qbook/internal/service"
)

func Run(ctx context.Context) error {

	// Инициализация репозиториев
	repo := memory.Init()

	// Инициализация сервисов
	ss := &service.Quotes{
		QuoteRepo:  repo,
		Identifier: &service.Identifier{},
	}

	// Инициализация и Запуск http контроллера
	server := initHttpServer(ss)

	ch := make(chan error)

	go func() {
		err := runHttpServer(ctx, server)
		ch <- err
	}()

	return <-ch
}
