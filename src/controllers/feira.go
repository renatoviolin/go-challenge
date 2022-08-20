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
	CreateFeiraService create_feira.CreateFeiraService
	UpdateFeiraService update_feira.UpdateFeiraService
	DeleteFeiraService delete_feira.DeleteFeiraService
	GetFeiraService    get_feira.GetFeiraService
}

func NewFeiraController(
	createFeira create_feira.CreateFeiraService,
	updateFeira update_feira.UpdateFeiraService,
	deleteFeira delete_feira.DeleteFeiraService,
	getFeira get_feira.GetFeiraService) FeiraController {
	return FeiraController{
		CreateFeiraService: createFeira,
		UpdateFeiraService: updateFeira,
		DeleteFeiraService: deleteFeira,
		GetFeiraService:    getFeira,
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
	return h.CreateFeiraService.Execute(feira)
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
	return h.UpdateFeiraService.Execute(feira)
}

func (h *FeiraController) Delete(id int64) error {
	return h.DeleteFeiraService.Execute(id)
}

func (h *FeiraController) FindByNome(query string) ([]dto.Feira, error) {
	return h.GetFeiraService.Execute(h.GetFeiraService.NomeFeira, query)
}

func (h *FeiraController) FindByBairro(query string) ([]dto.Feira, error) {
	return h.GetFeiraService.Execute(h.GetFeiraService.Bairro, query)
}

func (h *FeiraController) FindByDistrito(query string) ([]dto.Feira, error) {
	return h.GetFeiraService.Execute(h.GetFeiraService.Distrito, query)
}

func (h *FeiraController) FindByRegiao(query string) ([]dto.Feira, error) {
	return h.GetFeiraService.Execute(h.GetFeiraService.Regiao, query)
}
