package controllers

import (
	create_feira "api-unico/application/services/create-feira"
	delete_feira "api-unico/application/services/delete-feira"
	get_feira "api-unico/application/services/get-feira"
	update_feira "api-unico/application/services/update-feira"
	"api-unico/infra/database"
	"api-unico/repository/postgres"
	test_utils "api-unico/test-utils"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const conn = "postgres://postgres:root@localhost:5432/db-unico?sslmode=disable"

func generateInjections() FeiraController {
	connection := database.NewPostgreSQLConnection(conn)
	repository := postgres.NewfeiraRepositorySQL(connection.Db)
	createFeiraService := create_feira.NewCreateFeiraService(&repository)
	updateFeiraService := update_feira.NewUpdateFeiraService(&repository)
	deleteFeiraService := delete_feira.NewDeleteFeiraService(&repository)
	getFeiraService := get_feira.NewGetFeiraService(&repository)

	return FeiraController{
		CreateFeiraService: createFeiraService,
		UpdateFeiraService: updateFeiraService,
		DeleteFeiraService: deleteFeiraService,
		GetFeiraService:    getFeiraService,
	}

}

func Test_Create(t *testing.T) {
	feiraController := generateInjections()
	feiraPayload := test_utils.GenerateFeiraDto("nome da feira a partir do controller")
	id, err := feiraController.Create(feiraPayload)
	require.NoError(t, err)
	require.GreaterOrEqual(t, id, int64(1))
}

func Test_Update(t *testing.T) {
	feiraController := generateInjections()
	prevName := uuid.New().String()
	newName := uuid.New().String()
	feiraPayload := test_utils.GenerateFeiraDto(prevName)
	id, err := feiraController.Create(feiraPayload)
	require.NoError(t, err)

	feiraPayload.Id = id
	feiraPayload.NomeFeira = newName
	err = feiraController.Update(feiraPayload)
	require.NoError(t, err)

	_, err = feiraController.FindByNome(prevName)
	require.Error(t, err)

	validFeira, err := feiraController.FindByNome(newName)
	require.NoError(t, err)
	require.Equal(t, newName, validFeira[0].NomeFeira)
	require.Equal(t, id, validFeira[0].Id)
}

func Test_Delete(t *testing.T) {
	feiraController := generateInjections()
	name := uuid.New().String()
	feiraPayload := test_utils.GenerateFeiraDto(name)
	id, err := feiraController.Create(feiraPayload)
	require.NoError(t, err)
	require.GreaterOrEqual(t, id, int64(1))

	err = feiraController.Delete(id)
	require.NoError(t, err)

	_, err = feiraController.FindByNome(name)
	require.Error(t, err)
}

func Test_FindBy(t *testing.T) {
	feiraController := generateInjections()
	nomeFeira := uuid.New().String()
	regiao := nomeFeira + "-regiao"
	bairro := nomeFeira + "-bairro"
	distrito := nomeFeira + "-distrito"

	feiraPayload := test_utils.GenerateFeiraDto(nomeFeira)
	feiraPayload.NomeFeira = nomeFeira
	feiraPayload.Regiao5 = regiao
	feiraPayload.Distrito = distrito
	feiraPayload.Bairro = bairro

	id, err := feiraController.Create(feiraPayload)
	require.NoError(t, err)
	require.GreaterOrEqual(t, id, int64(1))

	result, err := feiraController.FindByNome(nomeFeira)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)

	result, err = feiraController.FindByBairro(bairro)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)

	result, err = feiraController.FindByDistrito(distrito)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)

	result, err = feiraController.FindByRegiao(regiao)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)
}
