package main

import (
	create_feira "api-unico/application/services/create-feira"
	delete_feira "api-unico/application/services/delete-feira"
	get_feira "api-unico/application/services/get-feira"
	update_feira "api-unico/application/services/update-feira"
	"api-unico/controllers"
	"api-unico/infra/database"
	http_server "api-unico/infra/http-server"
	"api-unico/infra/logger"
	"api-unico/repository/postgres"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

func loadVars() {
	err1 := gotenv.Load(".env")
	err2 := gotenv.Load("../.env")
	if err1 != nil && err2 != nil {
		logger.LogFatal("load-vars", err1.Error())
	}
}

func generateInjections() controllers.FeiraController {
	loadVars()
	fmt.Println(os.Getenv("POSTGRES_CONNECTION_STRING"))
	connection := database.NewPostgreSQLConnection(os.Getenv("POSTGRES_CONNECTION_STRING"))
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

func main() {
	logger.InitLog()

	feiraController := generateInjections()
	router := http_server.NewHttpServer(feiraController)
	port := 8000
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Logger.Info().Msgf("Starting server on port %d", port)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
