package app

import (
	"context"

	"github.com/Muhammed19m/qbook/internal/repository/memory"
)

func Run(ctx context.Context) error {
	generalConfig, err := InitConfig()
	if err != nil {
		return err
	}

	// Инициализация репозиториев
	repo := memory.Init()

	// Инициализация сервисов
	ss := initServices(repo)
	// Инициализация и Запуск http контроллера
	server := initHttpServer(&ss, generalConfig.HttpServer)

	return runHttpServer(ctx, server)
}
