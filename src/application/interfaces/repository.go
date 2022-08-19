package interfaces

import "api-unico/dto"

type FeiraRepository interface {
	Create(dto.Feira) error
	Update(dto.Feira) error
}
