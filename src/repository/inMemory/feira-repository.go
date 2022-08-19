package inMemory

import (
	"api-unico/dto"
	"errors"
	"math/rand"
)

type feiraRepository struct {
	database []dto.Feira
}

func NewFeiraRepository() feiraRepository {
	return feiraRepository{}
}

func (h *feiraRepository) Get(id int64) (dto.Feira, error) {
	for _, el := range h.database {
		if el.Id == id {
			return el, nil
		}
	}
	return dto.Feira{}, errors.New("resource not found")
}

func (h *feiraRepository) Create(feira dto.Feira) (int64, error) {
	id := rand.Int63n(1000000000)
	feira.Id = id
	h.database = append(h.database, feira)
	return id, nil
}

func (h *feiraRepository) Update(feira dto.Feira) error {
	for i, el := range h.database {
		if el.Id == feira.Id {
			h.database[i] = feira
			return nil
		}
	}
	return errors.New("resource not found")
}

func (h *feiraRepository) Delete(id int64) error {
	for i, el := range h.database {
		if el.Id == id {
			h.database = append(h.database[:i], h.database[i+1:]...)
			return nil
		}
	}
	return errors.New("resource not found")
}
