package main

import (
	"fmt"
	"net/http"

	"github.com/akshaym-3255/go-go-learning/crud-go/handler"
	"github.com/akshaym-3255/go-go-learning/crud-go/repository"
	"github.com/akshaym-3255/go-go-learning/crud-go/service"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from home page 1")
}

func main() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/", home)

	repo := repository.NewAlbumRepo()

	fmt.Printf("%p\n", repo)
	service := service.NewAlbumService(repo)
	albumhandler := handler.NewAlbumHandler(service)

	handler.RegisterHandler(newRouter, albumhandler)
	http.ListenAndServe(":8080", newRouter)
}

// type Article struct {
// 	Id      string `json:"id"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Contnet string `json:"content"`
// }

// var Articles []Article

// func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var article Article
// 	json.Unmarshal(reqBody, &article)

// 	Articles = append(Articles, article)

// 	json.NewEncoder(w).Encode(article)
// }

// func GetArticles(w http.ResponseWriter, r *http.Request) {
// 	logger.Debugf("reqest recivied")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func GetSingleArticle(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Articles {
// 		if article.Id == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}

// }
// func handleRoutes() {
// 	newRouter := mux.NewRouter().StrictSlash(true)
// 	newRouter.HandleFunc("/", home)
// 	newRouter.HandleFunc("/articles", CreateNewArticle).Methods("POST")
// 	newRouter.HandleFunc("/articles", GetArticles)
// 	newRouter.HandleFunc("/articles/{id}", GetSingleArticle)
// 	http.ListenAndServe(":8080", newRouter)
// }
