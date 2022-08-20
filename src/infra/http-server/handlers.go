package http_server

import (
	"api-unico/errors"

	"github.com/gin-gonic/gin"
)

func (h *HttpServer) healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "online"})
}

func (h *HttpServer) findByNome(ctx *gin.Context) {
	queryParam := ctx.Params.ByName("query")
	if queryParam == "" {
		ctx.JSON(400, gin.H{"message": "invalid query parameter"})
		return
	}

	output, err := h.feiraController.GetFeiraService.Execute(h.feiraController.GetFeiraService.NomeFeira, queryParam)
	if err != nil {
		if err == errors.ErrResourceNotFound {
			ctx.JSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": output})
}

func (h *HttpServer) findByRegiao(ctx *gin.Context) {
	queryParam := ctx.Params.ByName("query")
	if queryParam == "" {
		ctx.JSON(400, gin.H{"message": "invalid query parameter"})
		return
	}

	output, err := h.feiraController.GetFeiraService.Execute(h.feiraController.GetFeiraService.Regiao, queryParam)
	if err != nil {
		if err == errors.ErrResourceNotFound {
			ctx.JSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": output})
}

func (h *HttpServer) findByBairro(ctx *gin.Context) {
	queryParam := ctx.Params.ByName("query")
	if queryParam == "" {
		ctx.JSON(400, gin.H{"message": "invalid query parameter"})
		return
	}

	output, err := h.feiraController.GetFeiraService.Execute(h.feiraController.GetFeiraService.Bairro, queryParam)
	if err != nil {
		if err == errors.ErrResourceNotFound {
			ctx.JSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": output})
}

func (h *HttpServer) findByDistrito(ctx *gin.Context) {
	queryParam := ctx.Params.ByName("query")
	if queryParam == "" {
		ctx.JSON(400, gin.H{"message": "invalid query parameter"})
		return
	}

	output, err := h.feiraController.GetFeiraService.Execute(h.feiraController.GetFeiraService.Distrito, queryParam)
	if err != nil {
		if err == errors.ErrResourceNotFound {
			ctx.JSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": output})
}
