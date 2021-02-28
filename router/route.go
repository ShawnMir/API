package router

import (
	"log"
	"net/http"
	"restfulapi/handler"

	"github.com/gorilla/mux"
)

// HandleRequests ...
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", handler.HomePage)
	myRouter.HandleFunc("/articles", handler.ReturnAllArticles)
	myRouter.HandleFunc("/article", handler.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", handler.ReturnSingleArticle)
	myRouter.HandleFunc("/article/{id}", handler.DeleteArticle).Methods("Delete")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
