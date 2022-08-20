package http_server

import (
	create_feira "api-unico/application/services/create-feira"
	delete_feira "api-unico/application/services/delete-feira"
	get_feira "api-unico/application/services/get-feira"
	update_feira "api-unico/application/services/update-feira"
	"api-unico/controllers"
	"api-unico/dto"
	"api-unico/infra/database"
	"api-unico/repository/postgres"
	test_utils "api-unico/test-utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type outputPayload struct {
	Data []dto.Feira `json:"data"`
}

const conn = "postgres://postgres:root@localhost:5432/db-unico?sslmode=disable"

func generateInjections() controllers.FeiraController {
	connection := database.NewPostgreSQLConnection(conn)
	repository := postgres.NewfeiraRepositorySQL(connection.Db)
	createFeiraService := create_feira.NewCreateFeiraService(&repository)
	updateFeiraService := update_feira.NewUpdateFeiraService(&repository)
	deleteFeiraService := delete_feira.NewDeleteFeiraService(&repository)
	getFeiraService := get_feira.NewGetFeiraService(&repository)

	return controllers.FeiraController{
		CreateFeiraService: createFeiraService,
		UpdateFeiraService: updateFeiraService,
		DeleteFeiraService: deleteFeiraService,
		GetFeiraService:    getFeiraService,
	}
}

func Test_HealthCheck(t *testing.T) {
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.GET("/", httpServer.healthCheck)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.Equal(t, `{"message":"online"}`, string(responseData))
	require.Equal(t, 200, w.Code)
}

func Test_Create(t *testing.T) {
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)

	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)

	// valid payload
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// invalid payload
	payload.NomeFeira = ""
	jsonValue, _ = json.Marshal(payload)
	req, _ = http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Update(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.PUT("/api/v1/feira", httpServer.update)
	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// Read the created id
	type updateResponse struct {
		Id int `json:"id"`
	}
	var res updateResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	// perform update in the previous object
	payload.Id = int64(res.Id)
	jsonValue, _ = json.Marshal(payload)
	req, _ = http.NewRequest("PUT", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusAccepted, w.Code)

	// invalid update in the previous object
	payload.Id = -1
	jsonValue, _ = json.Marshal(payload)
	req, _ = http.NewRequest("PUT", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Delete(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.DELETE("/api/v1/feira/:id", httpServer.delete)

	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// Read the created id
	type updateResponse struct {
		Id int `json:"id"`
	}
	var res updateResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	// perform delete in the previous object
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/v1/feira/%d", res.Id), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	// invalid delete in the previous object
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/v1/feira/%d", -1), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_FindByNome(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.GET("/api/v1/feira/nome/:query", httpServer.findByNome)

	query := "nome-de-teste2"
	payload := test_utils.GenerateFeiraDto(query)
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// query for new object
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/feira/nome/%s", query), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var res outputPayload
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(res.Data), 1)

	// invalid query
	req, _ = http.NewRequest("GET", "/api/v1/feira/nome/xxxxxx", nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func Test_FindByRegiao(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.GET("/api/v1/feira/regiao/:query", httpServer.findByRegiao)

	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// query for new object
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/feira/regiao/%s", payload.Regiao5), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var res outputPayload
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(res.Data), 1)

	// invalid query
	req, _ = http.NewRequest("GET", "/api/v1/feira/regiao/xxxxxx", nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func Test_FindByDistrito(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.GET("/api/v1/feira/distrito/:query", httpServer.findByDistrito)

	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// query for new object
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/feira/distrito/%s", payload.Distrito), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var res outputPayload
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(res.Data), 1)

	// invalid query
	req, _ = http.NewRequest("GET", "/api/v1/feira/distrito/xxxxxx", nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func Test_FindByBairro(t *testing.T) {
	// Create a new object
	controller := generateInjections()
	httpServer := NewHttpServer(controller)
	httpServer.router.POST("/api/v1/feira", httpServer.create)
	httpServer.router.GET("/api/v1/feira/bairro/:query", httpServer.findByBairro)

	payload := test_utils.GenerateFeiraDto("nome-de-teste2")
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/feira", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// query for new object
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/feira/bairro/%s", payload.Bairro), nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var res outputPayload
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(res.Data), 1)

	// invalid query
	req, _ = http.NewRequest("GET", "/api/v1/feira/bairro/xxxxxx", nil)
	w = httptest.NewRecorder()
	httpServer.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}
