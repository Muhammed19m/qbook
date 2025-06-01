package service

import (
	"errors"
	"math/rand"

	"github.com/Muhammed19m/qbook/internal/domain"
)

type Quotes struct {
	QuoteRepo  domain.QuoteRepository
	Identifier *Identifier
}

type AddQuoteInput struct {
	Author string
	Text   string
}

func (in AddQuoteInput) Validate() error {
	quote := domain.Quote{
		Author: in.Author,
		Text:   in.Text,
	}
	err := errors.Join(quote.ValidateID(), quote.ValidateAuthor(), quote.ValidateText())
	if err != nil {
		return err
	}

	return nil
}

// AddQuote добавляет новую цитату
func (q *Quotes) AddQuote(in AddQuoteInput) (domain.Quote, error) {
	// Валидировать цитату
	if err := in.Validate(); err != nil {
		return domain.Quote{}, err
	}

	// Сохранить цитату
	quote := domain.Quote{
		ID:     q.Identifier.GetID(),
		Author: in.Author,
		Text:   in.Text,
	}
	if err := q.QuoteRepo.Save(quote); err != nil {
		return domain.Quote{}, err
	}

	return quote, nil
}

// AllQuotes возвращает все цитаты
func (q *Quotes) AllQuotes() ([]domain.Quote, error) {
	qs, err := q.QuoteRepo.List(domain.QuotesFilter{})
	if err != nil {
		return nil, err
	}

	return qs, nil
}

// GetRandomQuote возвращает случайную цитату
func (q *Quotes) GetRandomQuote() (domain.Quote, error) {
	randId := rand.Intn(q.Identifier.CurrenID() + 1)
	if randId == 0 {
		randId = 1
	}
	qs, err := q.QuoteRepo.List(domain.QuotesFilter{ID: randId})
	if err != nil {
		return domain.Quote{}, err
	}
	if len(qs) != 1 {
		return domain.Quote{}, nil
	}

	return qs[0], nil
}

type QuotesInput struct {
	Author string
}

type QuotesOut struct {
	Quotes []domain.Quote
}

// Quotes получает цитаты с учетом фильтрации
func (q *Quotes) Quotes(in QuotesInput) (QuotesOut, error) {
	quote := domain.Quote{Author: in.Author}
	if err := quote.ValidateAuthor(); err != nil {
		return QuotesOut{}, err
	}

	filter := domain.QuotesFilter{Author: in.Author}
	quotes, err := q.QuoteRepo.List(filter)
	if err != nil {
		return QuotesOut{}, err
	}
	return QuotesOut{
		Quotes: quotes,
	}, nil
}

type DeleteQuoteInput struct {
	ID int
}

func (in DeleteQuoteInput) Validate() error {
	if in.ID <= 0 {
		return errors.New("incorrect ID")
	}

	return nil
}

// DeleteQuote удаляет цитаты по ID
func (q *Quotes) DeleteQuote(in DeleteQuoteInput) error {
	if err := in.Validate(); err != nil {
		return err
	}
	err := q.QuoteRepo.Delete(in.ID)
	if err != nil {
		return err
	}

	return nil
}
