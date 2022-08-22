package http_server

import (
	"api-unico/controllers"
	"api-unico/infra/logger"
	"net/http"

	_ "api-unico/infra/http-server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Golang Challenge
// @version         1.0
// @host      localhost:8000
// @BasePath  /api/v1
type HttpServer struct {
	router          *gin.Engine
	feiraController controllers.FeiraController
}

func NewHttpServer(feiraController controllers.FeiraController) HttpServer {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	return HttpServer{
		router:          r,
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

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return h.router
}
