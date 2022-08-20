package postgres

import (
	"api-unico/infra/database"
	test_utils "api-unico/test-utils"
	"testing"

	"github.com/stretchr/testify/require"
)

const conn = "postgres://postgres:root@localhost:5432/db-unico?sslmode=disable"

func Test_Create(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)
	feira := test_utils.GenerateFeiraDto("feira-teste-1")
	id, err := repositorySQL.Create(feira)
	require.NoError(t, err)
	require.Greater(t, id, int64(0))
}

func Test_Update(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-2")
	id, err := repositorySQL.Create(feira)
	require.NoError(t, err)

	existFeira, err := repositorySQL.Get(id)
	require.NoError(t, err)

	existFeira.NomeFeira = "novo nome da feira"
	err = repositorySQL.Update(existFeira)
	require.NoError(t, err)

	existFeira2, err := repositorySQL.Get(id)
	require.NoError(t, err)
	require.Equal(t, existFeira.NomeFeira, existFeira2.NomeFeira)
}

func Test_Update_Invalid(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-2")
	feira.Id = -123
	err := repositorySQL.Update(feira)
	require.Error(t, err)
}

func Test_Delete(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-delete")
	id, err := repositorySQL.Create(feira)
	require.NoError(t, err)

	err = repositorySQL.Delete(id)
	require.NoError(t, err)

	_, err = repositorySQL.Get(id)
	require.Error(t, err)
}

func Test_Delete_Invalid(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-delete")
	_, err := repositorySQL.Create(feira)
	require.NoError(t, err)

	err = repositorySQL.Delete(12345)
	require.Error(t, err)

}

func Test_FindBy_Nome(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-find")
	_, err := repositorySQL.Create(feira)
	require.NoError(t, err)

	existingFeiras, err := repositorySQL.FindBy("nome_feira", "feira-teste-find")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(existingFeiras), 1)
}

func Test_FindBy_Regiao5(t *testing.T) {
	postgres := database.NewPostgreSQLConnection(conn)
	repositorySQL := NewfeiraRepositorySQL(postgres.Db)

	feira := test_utils.GenerateFeiraDto("feira-teste-find")
	feira.Regiao5 = "teste de regiao 5"
	_, err := repositorySQL.Create(feira)
	require.NoError(t, err)

	existingFeiras, err := repositorySQL.FindBy("regiao5", "teste")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(existingFeiras), 1)
}
