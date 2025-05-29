package service

import (
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

}
func Test_Quotes_QuoteByAuthor(t *testing.T) {

}
func Test_Quotes_DeleteQuote(t *testing.T) {

}
