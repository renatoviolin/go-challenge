package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Feira_Valid(t *testing.T) {
	feira := NewFeira()
	feira.Id = 1
	feira.Long = 123
	feira.Lat = 456
	feira.SetCens = 123123
	feira.Areap = 123
	feira.CodDist = "CodDist"
	feira.Distrito = "Distrito"
	feira.CodSubPref = 789
	feira.SubPrefe = "SubPref"
	feira.Regiao5 = "Regiao5"
	feira.Regiao8 = "Regiao8"
	feira.NomeFeira = "NomeFeira"
	feira.Registro = "Registro"
	feira.Logradouro = "Logradouro"
	feira.Numero = "Numero"
	feira.Bairro = "Bairro"
	feira.Referencia = "Referencia"

	valid := feira.IsValid()
	require.NoError(t, valid)
}

func Test_Feira_Invalid_Nome(t *testing.T) {
	feira := NewFeira()
	feira.NomeFeira = "   "
	feira.Regiao5 = "Regiao5"
	feira.Bairro = "Bairro"
	feira.Distrito = "Distrito"
	feira.Logradouro = "Logradouro"

	err := feira.IsValid()
	require.EqualError(t, ErrEmptyNomeFeira, err.Error())
}

func Test_Feira_Invalid_Regiao(t *testing.T) {
	feira := NewFeira()
	feira.NomeFeira = " aaaa  "
	feira.Regiao5 = " "
	feira.Bairro = "Bairro"
	feira.Distrito = "Distrito"
	feira.Logradouro = "Logradouro"

	err := feira.IsValid()
	require.EqualError(t, ErrEmptyRegiao5, err.Error())
}

func Test_Feira_Invalid_Bairro(t *testing.T) {
	feira := NewFeira()
	feira.NomeFeira = " aaaa  "
	feira.Regiao5 = "aaaa "
	feira.Bairro = ""
	feira.Distrito = "Distrito"
	feira.Logradouro = "Logradouro"

	err := feira.IsValid()
	require.EqualError(t, ErrEmptyBairro, err.Error())
}

func Test_Feira_Invalid_Distrito(t *testing.T) {
	feira := NewFeira()
	feira.NomeFeira = " aaaa  "
	feira.Regiao5 = "aaaa "
	feira.Bairro = "Bairro"
	feira.Distrito = ""
	feira.Logradouro = "Logradouro"

	err := feira.IsValid()
	require.EqualError(t, ErrEmptyDistrito, err.Error())
}

func Test_Feira_Invalid_Logradouro(t *testing.T) {
	feira := NewFeira()
	feira.NomeFeira = " aaaa  "
	feira.Regiao5 = "aaaa "
	feira.Bairro = "bairo"
	feira.Distrito = "Distrito"
	feira.Logradouro = ""

	err := feira.IsValid()
	require.EqualError(t, ErrEmptyLogradouro, err.Error())
}
