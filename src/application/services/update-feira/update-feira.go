package update_feira

import (
	"api-unico/application/entities"
	"api-unico/application/interfaces"
	"api-unico/dto"
	"api-unico/infra/logger"
)

type UpdateFeiraService struct {
	repository interfaces.FeiraRepository
}

func NewUpdateFeiraService(repository interfaces.FeiraRepository) UpdateFeiraService {
	return UpdateFeiraService{
		repository: repository,
	}
}

func (h *UpdateFeiraService) Execute(input entities.Feira) (err error) {
	if err := input.IsValid(); err != nil {
		logger.LogError("UpdateFeiraService", err.Error())
		return err
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
	if err = h.repository.Update(feiraPayload); err != nil {
		logger.LogError("UpdateFeiraService", err.Error())
		return err
	}

	return nil
}
