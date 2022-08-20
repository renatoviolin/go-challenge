package http_server

import (
	"github.com/gin-gonic/gin"
)

func (h *HttpServer) healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "online"})
}
