package delete_feira

import (
	create_feira "api-unico/application/services/create-feira"
	"api-unico/repository/inMemory"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DeleteFeira_Valid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	newId, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceDelete := NewDeleteFeiraService(&repository)
	err = serviceDelete.Execute(newId)
	require.NoError(t, err)

	_, err = repository.Get(newId)
	require.Error(t, err)
}

func Test_DeleteFeira_Invalid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome-feira")
	repository := inMemory.NewFeiraRepository()

	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	newId, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	serviceDelete := NewDeleteFeiraService(&repository)
	err = serviceDelete.Execute(int64(1234))
	require.Error(t, err)

	result, err := repository.Get(newId)
	require.NoError(t, err)
	require.Equal(t, feira.NomeFeira, result.NomeFeira)
}
