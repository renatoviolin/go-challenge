package http_server

import (
	"api-unico/dto"
	"api-unico/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary      Show server status
// @Success      200
// @Router       /health [get]
func (h *HttpServer) healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "online"})
}

// FindByName godoc
// @Summary      Find feira by nome
// @Produce      json
// @Param        query   path      string  true  "query"
// @Router       /feira/nome/{query} [get]
// @Success      200  {object}  dto.Feira
// @Failure      400
// @Failure      404
// @Failure      500
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

// FindByRegiao godoc
// @Summary      Find feira by regiao
// @Produce      json
// @Param        query   path      string  true  "query"
// @Router       /feira/regiao/{query} [get]
// @Success      200  {object}  dto.Feira
// @Failure      400
// @Failure      404
// @Failure      500
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

// FindByBairro godoc
// @Summary      Find feira by bairro
// @Produce      json
// @Param        query   path      string  true  "query"
// @Router       /feira/bairro/{query} [get]
// @Success      200  {object}  dto.Feira
// @Failure      400
// @Failure      404
// @Failure      500
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

// FindByDistrito godoc
// @Summary      Find feira by distrito
// @Produce      json
// @Param        query   path      string  true  "query"
// @Router       /feira/distrito/{query} [get]
// @Success      200  {object}  dto.Feira
// @Failure      400
// @Failure      404
// @Failure      500
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

// Create godoc
// @Summary      Create feira
// @Produce      json
// @Param        feira  body      dto.Feira  true  "Add feira"
// @Router       /feira [post]
// @Success      201
// @Failure      400
func (h *HttpServer) create(ctx *gin.Context) {
	var input dto.Feira
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	id, err := h.feiraController.Create(input)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": id})
}

// Update godoc
// @Summary      Update feira
// @Produce      json
// @Param        feira  body      dto.Feira  true  "Update feira"
// @Router       /feira [put]
// @Success      201
// @Failure      400
func (h *HttpServer) update(ctx *gin.Context) {
	var input dto.Feira
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = h.feiraController.Update(input)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(202, gin.H{"message": "successfully updated"})
}

// FindByDistrito godoc
// @Summary      Delete feira
// @Produce      json
// @Param        id   path      string  true  "id"
// @Router       /feira/{id} [delete]
// @Success      200
// @Failure      400
func (h *HttpServer) delete(ctx *gin.Context) {
	idString := ctx.Params.ByName("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid id parameter"})
		return
	}

	err = h.feiraController.Delete(id)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "successfully deleted"})
}
