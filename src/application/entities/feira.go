package entities

import (
	"api-unico/errors"
	"strings"
)

type Feira struct {
	CodDist    string
	Distrito   string
	SubPrefe   string
	Regiao5    string
	Regiao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
	SetCens    int64
	Id         int64
	Long       int64
	Lat        int64
	Areap      int64
	CodSubPref int64
}

func NewFeira() Feira {
	return Feira{}
}

func (h *Feira) IsValid() error {
	if len(strings.TrimSpace(h.NomeFeira)) == 0 {
		return errors.ErrEmptyNomeFeira
	}
	if len(strings.TrimSpace(h.Regiao5)) == 0 {
		return errors.ErrEmptyRegiao5
	}
	if len(strings.TrimSpace(h.Bairro)) == 0 {
		return errors.ErrEmptyBairro
	}
	if len(strings.TrimSpace(h.Distrito)) == 0 {
		return errors.ErrEmptyDistrito
	}
	if len(strings.TrimSpace(h.Logradouro)) == 0 {
		return errors.ErrEmptyLogradouro
	}
	return nil
}
