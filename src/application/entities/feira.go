package entities

import (
	"errors"
	"strings"
)

var (
	errEmptyNomeFeira  = errors.New("nome da feira vazio")
	errEmptyRegiao5    = errors.New("região5 vazio")
	errEmptyBairro     = errors.New("bairro vazio")
	errEmptyDistrito   = errors.New("distrito vazio")
	errEmptyLogradouro = errors.New("logradouro vazio")
)

type Feira struct {
	Id         int64
	Long       int64
	Lat        int64
	SetCens    string
	Areap      int64
	CodDist    string
	Distrito   string
	CodSubPref int64
	SubPrefe   string
	Regiao5    string
	Regiao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
}

func NewFeira() Feira {
	return Feira{}
}

func (h *Feira) isValid() error {
	if len(strings.TrimSpace(h.NomeFeira)) == 0 {
		return errEmptyNomeFeira
	}
	if len(strings.TrimSpace(h.Regiao5)) == 0 {
		return errEmptyRegiao5
	}
	if len(strings.TrimSpace(h.Bairro)) == 0 {
		return errEmptyBairro
	}
	if len(strings.TrimSpace(h.Distrito)) == 0 {
		return errEmptyDistrito
	}
	if len(strings.TrimSpace(h.Logradouro)) == 0 {
		return errEmptyLogradouro
	}
	return nil
}
