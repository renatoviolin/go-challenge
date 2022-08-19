package update_feira

import (
	create_feira "api-unico/application/services/create-feira"
	"api-unico/repository/inMemory"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UpdateFeira_Valid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome da feira")
	repository := inMemory.NewFeiraRepository()

	// create new feira
	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	newId, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	// update feira
	service := NewUpdateFeiraService(&repository)
	feira.Id = newId
	feira.NomeFeira = "novo nome de feira"
	err = service.Execute(feira)
	require.NoError(t, err)

	// retrieve feira
	feiraDto, err := repository.Get(newId)
	require.NoError(t, err)
	require.Equal(t, feira.NomeFeira, feiraDto.NomeFeira)
}

func Test_UpdateFeira_Invalid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome da feira")
	repository := inMemory.NewFeiraRepository()

	// create new feira
	serviceCreate := create_feira.NewCreateFeiraService(&repository)
	newId, err := serviceCreate.Execute(feira)
	require.NoError(t, err)

	// update feira
	service := NewUpdateFeiraService(&repository)
	feira.Id = int64(-1)
	feira.NomeFeira = "novo nome de feira"
	err = service.Execute(feira)
	require.Error(t, err)

	// retrieve feira
	feiraDto, err := repository.Get(newId)
	require.NoError(t, err)
	require.Equal(t, "nome da feira", feiraDto.NomeFeira)
}
