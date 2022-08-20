package inMemory

import (
	"api-unico/dto"
	"errors"
	"fmt"
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

func (h *feiraRepository) FindBy(column string, query string) ([]dto.Feira, error) {
	var result []dto.Feira
	if column == "distrito" {
		for i := range h.database {
			if h.database[i].Distrito == query {
				result = append(result, h.database[i])
			}
		}
		if len(result) == 0 {
			return result, errors.New("resource not found")
		}
		return result, nil
	}
	if column == "regiao5" {
		for i := range h.database {
			if h.database[i].Regiao5 == query {
				result = append(result, h.database[i])
			}
		}
		if len(result) == 0 {
			return result, errors.New("resource not found")
		}
		return result, nil
	}
	if column == "nome_feira" {
		for i := range h.database {
			if h.database[i].NomeFeira == query {
				result = append(result, h.database[i])
			}
		}
		if len(result) == 0 {
			return result, errors.New("resource not found")
		}
		return result, nil
	}
	if column == "bairro" {
		for i := range h.database {
			if h.database[i].Bairro == query {
				result = append(result, h.database[i])
			}
		}
		if len(result) == 0 {
			return result, errors.New("resource not found")
		}
		return result, nil
	}
	return result, fmt.Errorf("invalid query param: %s", column)
}
