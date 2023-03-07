package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Message string
}

func PingHandler(ctx *gin.Context) {
	fmt.Println()
	res := Resp{
		Message: "Pong",
	}
	fmt.Println(res)
	ctx.JSON(http.StatusOK, res)
}
