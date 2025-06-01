package app

import (
	"context"

	"github.com/Muhammed19m/qbook/internal/repository/memory"
)

func Run(ctx context.Context) error {

	// Инициализация репозиториев
	repo := memory.Init()

	// Инициализация сервисов
	ss := initServices(repo)
	// Инициализация и Запуск http контроллера
	server := initHttpServer(&ss)

	ch := make(chan error)

	go func() {
		err := runHttpServer(ctx, server)
		ch <- err
	}()

	return <-ch
}
