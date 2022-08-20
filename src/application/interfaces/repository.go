package interfaces

import "api-unico/dto"

type FeiraRepository interface {
	Create(dto.Feira) (int64, error)
	Update(dto.Feira) error
	Delete(id int64) error
	FindBy(column string, query string) ([]dto.Feira, error)
}
