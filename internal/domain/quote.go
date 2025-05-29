package domain

type Quote struct {
	ID     int
	Author string
	Text   string
}

type QuoteRepository interface {
	Save(quote Quote) error
	List(filter QuotesFilter) ([]Quote, error)
	Delete(id int) error
}

type QuotesFilter struct {
	ID     int
	Author string
}
