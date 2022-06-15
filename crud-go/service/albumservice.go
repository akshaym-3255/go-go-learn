package service

import (
	"fmt"

	"github.com/akshaym-3255/go-go-learning/crud-go/models"
	"github.com/akshaym-3255/go-go-learning/crud-go/repository"
)

type AlbumService interface {
	Get(id string) (models.Article, error)
	GetAll() ([]models.Article, error)
	Create(models.Article) error
}

type albumservice struct {
	repo repository.AlbumRepo
}

func NewAlbumService(repo repository.AlbumRepo) AlbumService {
	fmt.Printf("%p\n", repo)
	return &albumservice{repo}
}

func (service *albumservice) Get(id string) (models.Article, error) {
	article, err := service.repo.GetAlbum(id)

	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (service *albumservice) GetAll() ([]models.Article, error) {

	articles, err := service.repo.GetAll()

	// go asyncFunc()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (service *albumservice) Create(article models.Article) error {

	err := service.repo.Save(article)

	if err != nil {
		return err
	}

	return nil
}

// func asyncFunc() {
// 	for i := 0; i < 100; i++ {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println("doing processing")
// 		fmt.Println(i)

// 	}
// }
