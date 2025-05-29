package service

import "sync"

// Генератор ID
type Identifier struct {
	mu sync.Mutex
	id int
}

func (i *Identifier) GetID() int {
	var id int
	i.mu.Lock()
	i.id += 1
	id = i.id
	i.mu.Unlock()
	return id
}

func (i *Identifier) CurrenID() int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.id
}
