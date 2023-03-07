package router

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/ping", handler.PingHandler)
	r.GET("/albums", handler.GetAlbumsHandler)
}
