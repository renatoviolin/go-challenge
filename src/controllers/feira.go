package controllers

import (
	"api-unico/application/entities"
	create_feira "api-unico/application/services/create-feira"
	delete_feira "api-unico/application/services/delete-feira"
	get_feira "api-unico/application/services/get-feira"
	update_feira "api-unico/application/services/update-feira"
	"api-unico/dto"
)

type FeiraController struct {
	createFeiraService create_feira.CreateFeiraService
	updateFeiraService update_feira.UpdateFeiraService
	deleteFeiraService delete_feira.DeleteFeiraService
	getFeiraService    get_feira.GetFeiraService
}

func NewFeiraController(
	createFeira create_feira.CreateFeiraService,
	updateFeira update_feira.UpdateFeiraService,
	deleteFeira delete_feira.DeleteFeiraService,
	getFeira get_feira.GetFeiraService) FeiraController {
	return FeiraController{
		createFeiraService: createFeira,
		updateFeiraService: updateFeira,
		deleteFeiraService: deleteFeira,
		getFeiraService:    getFeira,
	}
}

func (h *FeiraController) Create(input dto.Feira) (int64, error) {
	feira := entities.Feira{
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
		SetCens:    input.SetCens,
		Id:         input.Id,
		Long:       input.Long,
		Lat:        input.Lat,
		Areap:      input.Areap,
		CodSubPref: input.CodSubPref,
	}
	return h.createFeiraService.Execute(feira)
}

func (h *FeiraController) Update(input dto.Feira) error {
	feira := entities.Feira{
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
		SetCens:    input.SetCens,
		Id:         input.Id,
		Long:       input.Long,
		Lat:        input.Lat,
		Areap:      input.Areap,
		CodSubPref: input.CodSubPref,
	}
	return h.updateFeiraService.Execute(feira)
}

func (h *FeiraController) Delete(id int64) error {
	return h.deleteFeiraService.Execute(id)
}

func (h *FeiraController) FindByNome(query string) ([]dto.Feira, error) {
	return h.getFeiraService.Execute(h.getFeiraService.NomeFeira, query)
}

func (h *FeiraController) FindByBairro(query string) ([]dto.Feira, error) {
	return h.getFeiraService.Execute(h.getFeiraService.Bairro, query)
}

func (h *FeiraController) FindByDistrito(query string) ([]dto.Feira, error) {
	return h.getFeiraService.Execute(h.getFeiraService.Distrito, query)
}

func (h *FeiraController) FindByRegiao(query string) ([]dto.Feira, error) {
	return h.getFeiraService.Execute(h.getFeiraService.Regiao, query)
}
