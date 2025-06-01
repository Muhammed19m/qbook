package memory

import (
	"errors"

	"github.com/Muhammed19m/qbook/internal/domain"
)

var (
	ErrQuoteNotCorrect = errors.New("quote not correct")
)

// Save запись задачи в хранилище
func (m *Memory) Save(quote domain.Quote) error {
	if quote.ID <= 0 {
		return ErrQuoteNotCorrect
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, quoteRepo := range m.quotes {
		if quote.ID == quoteRepo.ID {
			m.quotes[id] = quote
			return nil
		}
	}

	m.quotes = append(m.quotes, quote)

	return nil
}

func (m *Memory) List(filter domain.QuotesFilter) ([]domain.Quote, error) {
	var resultQuotes []domain.Quote
	m.mu.RLock()
	if filter.Author == "" && filter.ID == 0 {
		resultQuotes = make([]domain.Quote, len(m.quotes))
		copy(resultQuotes, m.quotes)
	} else if filter.Author == "" && filter.ID != 0 {
		for _, quote := range m.quotes {
			if quote.ID == filter.ID {
				resultQuotes = append(resultQuotes, quote)
				break // предпологаем что quote с таким ID всего 1
			}
		}
	} else if filter.Author != "" && filter.ID == 0 {
		for _, quote := range m.quotes {
			if quote.Author == filter.Author {
				resultQuotes = append(resultQuotes, quote)
			}
		}
	} else if filter.Author != "" && filter.ID != 0 {
		for _, quote := range m.quotes {
			if quote.Author == filter.Author && quote.ID == filter.ID {
				resultQuotes = append(resultQuotes, quote)
			}
		}
	}

	m.mu.RUnlock()
	return resultQuotes, nil
}

func (m *Memory) Delete(ID int) error {
	if ID <= 0 {
		return ErrQuoteNotCorrect
	}
	var idxCell int
	var ok bool
	m.mu.Lock()
	for idCell, quote := range m.quotes {
		if quote.ID == ID {
			idxCell = idCell
			ok = true
		}
	}
	if ok {
		lastQuote := m.quotes[len(m.quotes)-1]
		m.quotes[idxCell] = lastQuote
		m.quotes = m.quotes[:len(m.quotes)-1]
	}
	m.mu.Unlock()
	return nil
}
