package update_feira

import (
	"api-unico/application/entities"
	"api-unico/repository"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UpdateFeira_Valid(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("nome da feira")
	repository := repository.NewFeiraRepository()
	service := NewUpdateFeiraService(&repository)

	err := service.Execute(entities.Feira(feira))
	require.NoError(t, err)
}

func Test_UpdateFeira_Invalid_Name(t *testing.T) {
	feira := test_utils.GenerateFeiraEntity("  ")
	repository := repository.NewFeiraRepository()
	service := NewUpdateFeiraService(&repository)

	err := service.Execute(entities.Feira(feira))
	require.Equal(t, entities.ErrEmptyNomeFeira, err)
}
