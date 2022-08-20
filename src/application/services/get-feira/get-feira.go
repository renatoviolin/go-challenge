package get_feira

import (
	"api-unico/application/interfaces"
	"api-unico/dto"
	"api-unico/infra/logger"
)

type queryParam struct {
	value string
}

type GetFeiraService struct {
	repository interfaces.FeiraRepository
	Distrito   queryParam
	Regiao     queryParam
	NomeFeira  queryParam
	Bairro     queryParam
}

func NewGetFeiraService(repository interfaces.FeiraRepository) GetFeiraService {
	return GetFeiraService{
		repository: repository,
		Distrito:   queryParam{"distrito"},
		Regiao:     queryParam{"regiao5"},
		NomeFeira:  queryParam{"nome_feira"},
		Bairro:     queryParam{"bairro"},
	}
}

func (h *GetFeiraService) Execute(queryBy queryParam, queryString string) (output []dto.Feira, err error) {
	if output, err = h.repository.FindBy(queryBy.value, queryString); err != nil {
		logger.LogError("GetFeiraService", err.Error())
		return output, err
	}

	return output, nil
}

// distrito
// regiao5
// nome_feira
// bairro
