package repository

import (
	"api/models"
	"encoding/json"
	"fmt"
	"os"
)

type AlbumRepo interface {
	GetAlbums() ([]models.Album, error)
	AddAlbum() (models.Album, error)
}

type albumRepoer struct {
}

func NewAlbumRepo() AlbumRepo {
	return &albumRepoer{}
}

func (a *albumRepoer) GetAlbums() ([]models.Album, error) {
	var albums []models.Album
	content, err := os.ReadFile("repository/album.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &albums)
	if err != nil {
		return nil, err
	}
	fmt.Println(albums)
	return albums, nil
}

func (a *albumRepoer) AddAlbum() (models.Album, error) {
	return models.Album{}, nil
}
