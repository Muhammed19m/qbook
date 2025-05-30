package service

import (
	"slices"
	"testing"

	"github.com/Muhammed19m/qbook/internal/domain"
	assert "github.com/Muhammed19m/qbook/internal/domain/helpers_tests"
	"github.com/Muhammed19m/qbook/internal/repository/memory"
)

func NewServiceQuotes() *Quotes {
	return &Quotes{
		QuoteRepo:  memory.Init(),
		Identifier: &Identifier{},
	}
}

func Test_Quotes_AddQuote(t *testing.T) {
	t.Run("после добавления цитаты, цитата будет существовать в памяти", func(t *testing.T) {
		q := NewServiceQuotes()
		quote, err := q.AddQuote(AddQuoteInput{
			Author: "Author",
			Text:   "Quote",
		})
		assert.NoError(t, err)
		quoteRepo, err := q.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, 1)
		assert.Equal(t, quote, quoteRepo[0])
	})
	t.Run("можно добавлять много цитат", func(t *testing.T) {
		q := NewServiceQuotes()
		for range 10 {
			_, err := q.AddQuote(AddQuoteInput{
				Author: "Author",
				Text:   "Quote",
			})
			assert.NoError(t, err)
		}
		quoteRepo, err := q.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, 10)
	})
	t.Run("Author должен быть валдиен", func(t *testing.T) {
		q := NewServiceQuotes()
		_, err := q.AddQuote(AddQuoteInput{
			Author: "",
			Text:   "Quote",
		})
		assert.Error(t, err)

		quoteRepo, err := q.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, 0)
	})

	t.Run("Text не должен быть пустым", func(t *testing.T) {
		q := NewServiceQuotes()
		_, err := q.AddQuote(AddQuoteInput{
			Author: "Author",
			Text:   "",
		})
		assert.Error(t, err)

		quoteRepo, err := q.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, 0)
	})
}

func Test_Quotes_AllQuotes(t *testing.T) {
	t.Run("из пустой памяти, вернет пустой список", func(t *testing.T) {
		q := NewServiceQuotes()

		quoteRepo, err := q.AllQuotes()
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, 0)
	})
	t.Run("вернет все цитаты, которые существуют", func(t *testing.T) {
		q := NewServiceQuotes()
		const count = 15
		for range count {
			err := q.QuoteRepo.Save(domain.Quote{
				ID: q.Identifier.GetID(),
			})
			assert.NoError(t, err)
		}

		quoteRepo, err := q.AllQuotes()
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, count)
	})
	t.Run("удаленные цитаты не возвращает", func(t *testing.T) {
		q := NewServiceQuotes()
		const count = 15
		for range count {
			err := q.QuoteRepo.Save(domain.Quote{
				ID: q.Identifier.GetID(),
			})
			assert.NoError(t, err)
		}
		err := q.QuoteRepo.Delete(q.Identifier.CurrenID())
		assert.NoError(t, err)

		err = q.QuoteRepo.Delete(q.Identifier.CurrenID() - 1)
		assert.NoError(t, err)

		quoteRepo, err := q.AllQuotes()
		assert.NoError(t, err)
		assert.Len(t, quoteRepo, count-2)
	})
}

func Test_Quotes_GetRandomQuote(t *testing.T) {
	t.Run("если цитат не существует, то вернет zero value", func(t *testing.T) {
		q := NewServiceQuotes()
		quote, err := q.GetRandomQuote()
		assert.NoError(t, err)
		assert.Equal(t, domain.Quote{}, quote)
	})
	t.Run("если цитата одна, то каждый раз будет она возвращаться", func(t *testing.T) {
		q := NewServiceQuotes()
		quote := domain.Quote{
			ID:     q.Identifier.GetID(),
			Author: "Author",
			Text:   "Text",
		}
		err := q.QuoteRepo.Save(quote)
		assert.NoError(t, err)
		for range 100 {
			quoteRepo, err := q.GetRandomQuote()
			assert.NoError(t, err)
			assert.Equal(t, quoteRepo, quote)
		}
	})
	t.Run("если цитат несколько, то каждый раз случайный из них будет возвращаться", func(t *testing.T) {
		q := NewServiceQuotes()
		quotes := []domain.Quote{
			{
				ID:     q.Identifier.GetID(),
				Author: "Author1",
				Text:   "Text1",
			},
			{
				ID:     q.Identifier.GetID(),
				Author: "Author2",
				Text:   "Text2",
			},
			{
				ID:     q.Identifier.GetID(),
				Author: "Author3",
				Text:   "Text3",
			},
		}
		for _, quote := range quotes {
			err := q.QuoteRepo.Save(quote)
			assert.NoError(t, err)
		}
		assert.Equal(t, q.Identifier.CurrenID(), 3)

		qr, err := q.GetRandomQuote()
		assert.NoError(t, err)

		if !slices.Contains(quotes, qr) {
			t.Error()
		}
	})
}
func Test_Quotes_QuoteByAuthor(t *testing.T) {
	t.Run("из пустого списка - вернется пустой список", func(t *testing.T) {
		qService := NewServiceQuotes()
		qs, err := qService.QuoteByAuthor(QuoteByAuthorInput{
			Author: "Muhammed",
		})
		assert.NoError(t, err)
		assert.Len(t, qs, 0)
	})
	t.Run("передаваемый аргумент Author должен быть валиден", func(t *testing.T) {
		qService := NewServiceQuotes()
		qs, err := qService.QuoteByAuthor(QuoteByAuthorInput{
			Author: "",
		})
		assert.Error(t, err)
		assert.Len(t, qs, 0)
	})
	t.Run("если у этого Автора нету цитат, то вернется пустой список", func(t *testing.T) {
		qService := NewServiceQuotes()
		quotes := []domain.Quote{
			{
				ID:     qService.Identifier.GetID(),
				Author: "Author1",
				Text:   "Text1",
			},
			{
				ID:     qService.Identifier.GetID(),
				Author: "Author2",
				Text:   "Text2",
			},
			{
				ID:     qService.Identifier.GetID(),
				Author: "Author3",
				Text:   "Text3",
			},
		}
		for _, quote := range quotes {
			err := qService.QuoteRepo.Save(quote)
			assert.NoError(t, err)
		}

		qs, err := qService.QuoteByAuthor(QuoteByAuthorInput{
			Author: "Muhammed",
		})
		assert.NoError(t, err)
		assert.Len(t, qs, 0)
	})
	t.Run("если у этого Автора есть цитата, то вернется список только из его цитат", func(t *testing.T) {
		qService := NewServiceQuotes()
		quotes := []domain.Quote{
			{
				ID:     qService.Identifier.GetID(),
				Author: "Muhammed",
				Text:   "Text1",
			},
			{
				ID:     qService.Identifier.GetID(),
				Author: "Author2",
				Text:   "Text2",
			},
			{
				ID:     qService.Identifier.GetID(),
				Author: "Muhammed",
				Text:   "Text3",
			},
		}
		for _, quote := range quotes {
			err := qService.QuoteRepo.Save(quote)
			assert.NoError(t, err)
		}

		qs, err := qService.QuoteByAuthor(QuoteByAuthorInput{
			Author: "Muhammed",
		})
		assert.NoError(t, err)
		assert.Len(t, qs, 2)
		assert.Equal(t, quotes[0], qs[0])
		assert.Equal(t, quotes[2], qs[1])
	})
}
func Test_Quotes_DeleteQuote(t *testing.T) {
	t.Run("удаление не существующей цитаты - не вернет ошибку", func(t *testing.T) {
		qService := NewServiceQuotes()
		err := qService.QuoteRepo.Delete(1)
		assert.NoError(t, err)
	})
	t.Run("удаление не существующей цитаты - не тронет другие цитаты", func(t *testing.T) {
		qService := NewServiceQuotes()
		err := qService.QuoteRepo.Save(domain.Quote{ID: 1})
		assert.NoError(t, err)

		err = qService.QuoteRepo.Delete(2)
		assert.NoError(t, err)

		quotes, err := qService.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
	})
	t.Run("удаление существующей цитаты", func(t *testing.T) {
		qService := NewServiceQuotes()
		err := qService.QuoteRepo.Save(domain.Quote{ID: 1})
		assert.NoError(t, err)

		err = qService.QuoteRepo.Delete(1)
		assert.NoError(t, err)

		quotes, err := qService.QuoteRepo.List(domain.QuotesFilter{})
		assert.NoError(t, err)
		assert.Len(t, quotes, 0)
	})
}
