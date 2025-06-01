package repository_tests

import (
	"testing"

	"github.com/Muhammed19m/qbook/internal/domain"
	"github.com/Muhammed19m/qbook/pkg/assert"
)

func Test_QuoteRepository(t *testing.T, newRepository func() domain.QuoteRepository) {
	t.Run("Save", func(t *testing.T) {
		t.Run("сохранение quote толкьо с ID - можно", func(t *testing.T) {
			rep := newRepository()
			q := domain.Quote{ID: 1}
			err := rep.Save(q)
			assert.NoError(t, err)
		})
		t.Run("сохранение quote без ID - ошибка", func(t *testing.T) {
			rep := newRepository()
			q := domain.Quote{}
			err := rep.Save(q)
			assert.Error(t, err)
		})
		t.Run("полноценная quote - можно", func(t *testing.T) {
			rep := newRepository()
			q := domain.Quote{
				ID:     1,
				Author: "Автор",
				Text:   "прекрасное небо 123",
			}
			err := rep.Save(q)
			assert.NoError(t, err)
		})

		t.Run("У кажой quote, ID должен быть уникальным, иначе Save изменит quote с таким ID", func(t *testing.T) {
			rep := newRepository()
			q1 := domain.Quote{
				ID:     1,
				Author: "Автор",
				Text:   "прекрасное небо 123",
			}
			err := rep.Save(q1)
			assert.NoError(t, err)

			q2 := domain.Quote{
				ID:     1,
				Author: "Author",
				Text:   "Quote",
			}
			err = rep.Save(q2)
			assert.NoError(t, err)

			quotes, err := rep.List(domain.QuotesFilter{})
			assert.Len(t, quotes, 1)
			assert.NoError(t, err)
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("из пустого репозитория вернутся пустой список", func(t *testing.T) {
			rep := newRepository()
			quotes, err := rep.List(domain.QuotesFilter{})
			assert.Len(t, quotes, 0)
			assert.NoError(t, err)
		})

		t.Run("из пустого фильтра вернутся все цитаты", func(t *testing.T) {
			rep := newRepository()
			const countQuotes = 10
			for i := range countQuotes {
				err := rep.Save(domain.Quote{ID: i + 1})
				assert.NoError(t, err)
			}
			quotes, err := rep.List(domain.QuotesFilter{})
			assert.Len(t, quotes, countQuotes)
			assert.NoError(t, err)
		})
		t.Run("фильтр по ID", func(t *testing.T) {
			rep := newRepository()
			const countQuotes = 10
			for i := range countQuotes {
				err := rep.Save(domain.Quote{ID: i + 1})
				assert.NoError(t, err)
			}
			quotes, err := rep.List(domain.QuotesFilter{ID: 5})
			assert.Len(t, quotes, 1)
			assert.NoError(t, err)
			assert.Equal(t, 5, quotes[0].ID)
		})
		t.Run("пустой спиосок, если фильтр по ID который не существует", func(t *testing.T) {
			rep := newRepository()
			const countQuotes = 10
			for i := range countQuotes {
				err := rep.Save(domain.Quote{ID: i + 1})
				assert.NoError(t, err)
			}
			quotes, err := rep.List(domain.QuotesFilter{ID: 100})
			assert.Len(t, quotes, 0)
			assert.NoError(t, err)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("удаление не существующей quote не вернет ошибку", func(t *testing.T) {
			rep := newRepository()
			err := rep.Delete(1)
			assert.NoError(t, err)
		})
		t.Run("ID должен быть валиден", func(t *testing.T) {
			rep := newRepository()
			err := rep.Delete(0)
			assert.Error(t, err)
		})
		t.Run("если удалить quote, то она на удалится с памяти", func(t *testing.T) {
			rep := newRepository()
			err := rep.Save(domain.Quote{
				ID:     1,
				Author: "Author",
				Text:   "Quote",
			})
			assert.NoError(t, err)

			quotes, err := rep.List(domain.QuotesFilter{})
			assert.NoError(t, err)
			assert.Len(t, quotes, 1)

			err = rep.Delete(1)
			assert.NoError(t, err)

			quotes, err = rep.List(domain.QuotesFilter{})
			assert.NoError(t, err)
			assert.Len(t, quotes, 0)

		})
	})
}
