package service

import (
	"errors"
	"math/rand"

	"github.com/Muhammed19m/qbook/internal/domain"
)

type Quotes struct {
	QuoteRepo  domain.QuoteRepository
	Identifier Identifier
}

type AddQuoteInput struct {
	Author string
	Text   string
}

// добавление новой Цитаты
func (q *Quotes) AddQuote(in AddQuoteInput) (domain.Quote, error) {
	quote := domain.Quote{
		ID:     q.Identifier.GetID(),
		Author: in.Author,
		Text:   in.Text,
	}
	err := errors.Join(quote.ValidateID(), quote.ValidateAuthor(), quote.ValidateText())
	if err != nil {
		return domain.Quote{}, err
	}

	err = q.QuoteRepo.Save(quote)
	if err != nil {
		return domain.Quote{}, err
	}

	return quote, nil
}

// получение всех цитат
func (q *Quotes) AllQuotes() ([]domain.Quote, error) {
	qs, err := q.QuoteRepo.List(domain.QuotesFilter{})
	if err != nil {
		return nil, err
	}

	return qs, nil
}

// получение случайно цитаты
func (q *Quotes) GetRandomQuote() (domain.Quote, error) {
	randId := rand.Intn(q.Identifier.CurrenID()) + 1

	qs, err := q.QuoteRepo.List(domain.QuotesFilter{ID: randId})
	if err != nil {
		return domain.Quote{}, err
	}

	return qs[0], nil
}

type QuoteByAuthorInput struct {
	Author string
}

// получение цитат по фильтрации по автору
func (q *Quotes) QuoteByAuthor(in QuoteByAuthorInput) ([]domain.Quote, error) {
	quote := domain.Quote{Author: in.Author}
	if err := quote.ValidateAuthor(); err != nil {
		return nil, err
	}

	filter := domain.QuotesFilter{Author: in.Author}
	quotes, err := q.QuoteRepo.List(filter)
	if err != nil {
		return nil, err
	}
	return quotes, nil
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

// удаление цитаты по ID
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
