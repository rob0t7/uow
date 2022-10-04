package uow

import "github.com/google/uuid"

type Repository struct {
	entries map[uuid.UUID]Company
}

func (r *Repository) Insert(company *Company) error {
	if _, found := r.entries[company.ID]; found {
		return ErrConflict
	}
	r.entries[company.ID] = *company
	return nil
}

func (r *Repository) FindAll() ([]Company, error) {
	companies := make([]Company, 0, len(r.entries))
	for _, company := range r.entries {
		c := company
		companies = append(companies, c)
	}
	return companies, nil
}

func (r *Repository) FindByID(id uuid.UUID) (*Company, error) {
	company, found := r.entries[id]
	if found {
		c := company
		return &c, nil
	}
	return nil, ErrNotFound
}
