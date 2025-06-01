package app

import (
	"github.com/Muhammed19m/qbook/internal/repository/memory"
	"github.com/Muhammed19m/qbook/internal/service"
)

type services struct {
	quotes *service.Quotes
}

func (ss *services) Quotes() *service.Quotes {
	return ss.quotes
}

func initServices(repo *memory.Memory) services {
	return services{
		quotes: &service.Quotes{
			QuoteRepo:  repo,
			Identifier: &service.Identifier{},
		},
	}
}
