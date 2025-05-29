package memory

import (
	"testing"

	"github.com/Muhammed19m/qbook/internal/domain"
	"github.com/Muhammed19m/qbook/internal/domain/repository_tests"
)

func Test_QuoteRepository(t *testing.T) {
	repository_tests.Test_QuoteRepository(t, func() domain.QuoteRepository {
		return Init()
	})
}
