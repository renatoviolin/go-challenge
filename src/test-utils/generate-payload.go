package test_utils

import (
	"api-unico/application/entities"
	"api-unico/dto"
)

func GenerateFeiraEntity(nome string) entities.Feira {
	return entities.Feira{
		Long:       123,
		Lat:        456,
		SetCens:    "setCens",
		Areap:      123,
		CodDist:    "CodDist",
		Distrito:   "Distrito",
		CodSubPref: 789,
		SubPrefe:   "SubPref",
		Regiao5:    "Regiao5",
		Regiao8:    "Regiao8",
		NomeFeira:  nome,
		Registro:   "Registro",
		Logradouro: "Logradouro",
		Numero:     "Numero",
		Bairro:     "Bairro",
		Referencia: "Referencia",
	}
}

func GenerateFeiraDto(nome string) dto.Feira {
	return dto.Feira{
		Long:       123,
		Lat:        456,
		SetCens:    "setCens",
		Areap:      123,
		CodDist:    "CodDist",
		Distrito:   "Distrito",
		CodSubPref: 789,
		SubPrefe:   "SubPref",
		Regiao5:    "Regiao5",
		Regiao8:    "Regiao8",
		NomeFeira:  nome,
		Registro:   "Registro",
		Logradouro: "Logradouro",
		Numero:     "Numero",
		Bairro:     "Bairro",
		Referencia: "Referencia",
	}
}
