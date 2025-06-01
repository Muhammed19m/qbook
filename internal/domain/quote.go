package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Quote struct {
	ID     int
	Author string
	Text   string
}

func (q Quote) IsZero() bool {
	return q == (Quote{})
}

const AuthorNameMaxLen = 50

var (
	ErrQouteIDInvalid       = errors.New("некоректный QuoteID")
	ErrQouteAuthorInvalid   = errors.New("некоректный QuoteAuthor")
	ErrQouteTextInvalid     = errors.New("некоректный QuoteText")
	ErrEmptyAuthor          = errors.New("имя не может быть пустым")
	ErrMaxLenAuthor         = fmt.Errorf("длина имени не может превышать %d символов", AuthorNameMaxLen)
	ErrStartOrEndWithSpaces = errors.New("имя не может содержать начальных или конечных пробелов")
	ErrControlSymbol        = errors.New("имя не может содержать управляющих символов")
)

func (q Quote) ValidateID() error {
	if q.ID <= 0 {
		return ErrQouteIDInvalid
	}

	return nil
}

func (q Quote) ValidateAuthor() error {
	if strings.TrimSpace(q.Author) == "" {
		return ErrEmptyAuthor
	}

	if len([]rune(q.Author)) > AuthorNameMaxLen {
		return ErrMaxLenAuthor
	}

	if q.Author != strings.TrimSpace(q.Author) {
		return ErrStartOrEndWithSpaces
	}

	for _, r := range q.Author {
		if unicode.IsControl(r) && r != ' ' { // \n, \t, \r и т.д.
			return ErrControlSymbol
		}
	}

	return nil
}

func (q Quote) ValidateText() error {
	s := strings.TrimSpace(q.Text)
	if s == "" {
		return ErrQouteIDInvalid
	}

	return nil
}

type QuoteRepository interface {
	Save(quote Quote) error
	List(filter QuotesFilter) ([]Quote, error)
	//Random() (Quote, error)
	Delete(id int) error
}

type QuotesFilter struct {
	ID     int
	Author string
}
