package update_feira

import (
	"api-unico/application/entities"
	"api-unico/application/interfaces"
	"api-unico/dto"
)

type updateFeiraService struct {
	repository interfaces.FeiraRepository
}

func NewUpdateFeiraService(repository interfaces.FeiraRepository) updateFeiraService {
	return updateFeiraService{
		repository: repository,
	}
}

func (h *updateFeiraService) Execute(input entities.Feira) (err error) {
	if err := input.IsValid(); err != nil {
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
		return err
	}

	return nil
}
