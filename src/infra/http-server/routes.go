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
	h.router.GET("/api/v1/feira/nome/:query", h.findByNome)
	h.router.GET("/api/v1/feira/regiao/:query", h.findByRegiao)
	h.router.GET("/api/v1/feira/distrito/:query", h.findByDistrito)
	h.router.GET("/api/v1/feira/bairro/:query", h.findByBairro)

	h.router.POST("/api/v1/feira", h.create)
	h.router.PUT("/api/v1/feira", h.update)
	h.router.DELETE("/api/v1/feira/:id", h.delete)

	return h.router
}
