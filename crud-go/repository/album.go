package repository

import "github.com/akshaym-3255/go-go-learning/crud-go/models"

var Articles = []models.Article{
	{Id: "1", Title: "hello", Desc: "test", Contnet: "this is first article"},
	{Id: "2	", Title: "hello 2", Desc: "test", Contnet: "this is second Article"},
}

type AlbumRepo interface {
	GetAll() ([]models.Article, error)
	GetAlbum(id string) (models.Article, error)
	Save(models.Article) error
}

type albumrepo struct {
}

func NewAlbumRepo() AlbumRepo {
	return &albumrepo{}
}

func (repo *albumrepo) GetAll() ([]models.Article, error) {
	return Articles, nil
}

func (repo *albumrepo) GetAlbum(id string) (models.Article, error) {
	var article models.Article
	for _, CurrArticle := range Articles {
		if CurrArticle.Id == id {
			article = CurrArticle
		}
	}
	return article, nil
}

func (repo *albumrepo) Save(article models.Article) error {
	Articles = append(Articles, article)
	return nil
}
