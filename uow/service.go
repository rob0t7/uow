package uow

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrConflict = errors.New("conflict")
	ErrNotFound = errors.New("not found")
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	if repository == nil {
		panic("repository is required")
	}
	return &Service{repository: repository}
}

func (s *Service) RegisterCompany(name string) (*Company, error) {
	company := Company{
		ID:   uuid.New(),
		Name: name,
	}
	if err := s.repository.Insert(&company); err != nil {
		return nil, err
	}
	return &company, nil
}

func (s *Service) FindAll() ([]Company, error) {
	return s.repository.FindAll()
}

func (s *Service) FIndByID(id uuid.UUID) (*Company, error) {
	return s.repository.FindByID(id)
}
