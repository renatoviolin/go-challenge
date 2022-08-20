package get_feira

import (
	"api-unico/application/interfaces"
	"api-unico/dto"
)

type queryParam struct {
	value string
}

type getFeiraService struct {
	repository interfaces.FeiraRepository
	Distrito   queryParam
	Regiao     queryParam
	NomeFeira  queryParam
	Bairro     queryParam
}

func NewGetFeiraService(repository interfaces.FeiraRepository) getFeiraService {
	return getFeiraService{
		repository: repository,
		Distrito:   queryParam{"distrito"},
		Regiao:     queryParam{"regiao5"},
		NomeFeira:  queryParam{"nome_feira"},
		Bairro:     queryParam{"bairro"},
	}
}

func (h *getFeiraService) Execute(queryBy queryParam, queryString string) (output []dto.Feira, err error) {
	if output, err = h.repository.FindBy(queryBy.value, queryString); err != nil {
		return output, err
	}

	return output, nil
}

// distrito
// regiao5
// nome_feira
// bairro