package create_feira

import (
	"api-unico/application/entities"
	"api-unico/application/interfaces"
	"api-unico/dto"
)

type createFeiraService struct {
	repository interfaces.FeiraRepository
}

func NewCreateFeiraService(repository interfaces.FeiraRepository) createFeiraService {
	return createFeiraService{
		repository: repository,
	}
}

func (h *createFeiraService) Execute(input entities.Feira) (id int64, err error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	feiraPayload := dto.Feira{
		SetCens:    input.SetCens,
		CodDist:    input.CodDist,
		Distrito:   input.Distrito,
		SubPrefe:   input.SubPrefe,
		Regiao5:    input.Regiao5,
		Regiao8:    input.Regiao8,
		NomeFeira:  input.NomeFeira,
		Registro:   input.Registro,
		Logradouro: input.Logradouro,
		Numero:     input.Numero,
		Bairro:     input.Bairro,
		Referencia: input.Referencia,
		Id:         input.Id,
		Long:       input.Long,
		Lat:        input.Lat,
		Areap:      input.Areap,
		CodSubPref: input.CodSubPref,
	}
	id, err = h.repository.Create(feiraPayload)
	if err != nil {
		return 0, err
	}

	return id, err
}
