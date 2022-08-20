package get_feira

import (
	create_feira "api-unico/application/services/create-feira"
	"api-unico/repository/inMemory"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetFeira_ByNomeFeira(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	_, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceGetFeira := NewGetFeiraService(&repository)
	results, err := serviceGetFeira.Execute("nome-feira", serviceGetFeira.NomeFeira)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(results), 1)
}

func Test_GetFeira_ByRegiao5(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	feira.Regiao5 = "nome de uma regiao 5"
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	_, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceGetFeira := NewGetFeiraService(&repository)
	results, err := serviceGetFeira.Execute("nome de uma regiao 5", serviceGetFeira.Regiao)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(results), 1)
}

func Test_GetFeira_ByDistrito(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	feira.Distrito = "nome de um distrito"
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	_, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceGetFeira := NewGetFeiraService(&repository)
	results, err := serviceGetFeira.Execute("nome de um distrito", serviceGetFeira.Distrito)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(results), 1)
}

func Test_GetFeira_ByBairro(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	feira.Bairro = "nome de um bairro"
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	_, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceGetFeira := NewGetFeiraService(&repository)
	results, err := serviceGetFeira.Execute("nome de um bairro", serviceGetFeira.Bairro)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(results), 1)
}

func Test_GetFeira_NotFound(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	_, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceGetFeira := NewGetFeiraService(&repository)
	results, err := serviceGetFeira.Execute("nome-feira-2", serviceGetFeira.NomeFeira)
	require.Error(t, err)
	require.Empty(t, results)
}
