package http_server

import (
	"api-unico/controllers"
	"api-unico/infra/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	router          *gin.Engine
	feiraController controllers.FeiraController
}

func NewHttpServer(feiraController controllers.FeiraController) HttpServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	return HttpServer{
		router:          gin.Default(),
		feiraController: feiraController,
	}
}

func (h *HttpServer) Routes() http.Handler {
	h.router.Use(logger.HttpLogger())

	h.router.GET("/api/v1/health", h.healthCheck)
	return h.router
}
