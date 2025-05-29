package memory

import (
	"sync"

	"github.com/Muhammed19m/qbook/internal/domain"
)

type Memory struct {
	mu     sync.RWMutex
	quotes []domain.Quote
}

func Init() *Memory {
	return &Memory{}
}
