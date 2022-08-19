package repository

import "api-unico/dto"

type feiraRepository struct {
}

func NewFeiraRepository() feiraRepository {
	return feiraRepository{}
}

func (h *feiraRepository) Create(dto.Feira) error {
	return nil
}

func (h *feiraRepository) Update(dto.Feira) error {
	return nil
}
