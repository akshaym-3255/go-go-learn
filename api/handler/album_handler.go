package handler

import (
	"api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbumsHandler(ctx *gin.Context) {
	repo := repository.NewAlbumRepo()
	albums, err := repo.GetAlbums()
	query := ctx.Query()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("new-id", "2")
	ctx.JSON(http.StatusOK, albums)
}

func ProcessAlbums(ctx *gin.Context) {
	
}
