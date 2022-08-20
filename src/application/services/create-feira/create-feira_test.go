package create_feira

import (
	"api-unico/application/entities"
	"api-unico/errors"
	"api-unico/repository/inMemory"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreateFeira_Valid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome da feira")
	repository := inMemory.NewFeiraRepository()
	service := NewCreateFeiraService(&repository)

	_, err := service.Execute(entities.Feira(feira))
	require.NoError(t, err)
}

func Test_CreateFeira_Invalid_Name(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("")
	repository := inMemory.NewFeiraRepository()
	service := NewCreateFeiraService(&repository)

	id, err := service.Execute(entities.Feira(feira))
	require.Equal(t, errors.ErrEmptyNomeFeira, err)
	require.Equal(t, id, int64(0))
}
