package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/akshaym-3255/go-go-learning/crud-go/models"
	"github.com/akshaym-3255/go-go-learning/crud-go/service"
	"github.com/gorilla/mux"
)

func RegisterHandler(router *mux.Router, albumHandler AlbumHandler) {
	router.HandleFunc("/articles", albumHandler.CreateNewArticle).Methods("POST")
	router.HandleFunc("/articles", albumHandler.GetArticles)
	router.HandleFunc("/articles/{id}", albumHandler.GetSingleArticle)
}

type AlbumHandler interface {
	CreateNewArticle(w http.ResponseWriter, r *http.Request)
	GetArticles(w http.ResponseWriter, r *http.Request)
	GetSingleArticle(w http.ResponseWriter, r *http.Request)
}
type albumHandler struct {
	service service.AlbumService
}

func NewAlbumHandler(service service.AlbumService) AlbumHandler {
	return &albumHandler{service}
}

func (handler *albumHandler) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article models.Article
	json.Unmarshal(reqBody, &article)

	err := handler.service.Create(article)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
	}

	json.NewEncoder(w).Encode(article)

}

func (handler *albumHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article

	articles, err := handler.service.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(articles)
	}
	json.NewEncoder(w).Encode(articles)
}

func (handler *albumHandler) GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var article models.Article

	article, err := handler.service.Get(id)
	if err != nil {
		json.NewEncoder(w).Encode(article)
	}
	json.NewEncoder(w).Encode(article)
}
