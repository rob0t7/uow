package uow

import (
	"github.com/google/uuid"
)

type Company struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Store struct {
	entries map[uuid.UUID]Company
}

func NewStore() *Store {
	return &Store{entries: make(map[uuid.UUID]Company)}
}
