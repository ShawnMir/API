package main

import (
	"fmt"

	"restfulapi/handler"
	"restfulapi/router"
)

func main() {
	fmt.Println("Creating Restful API using Mux Routers")

	handler.Articles = []handler.Article{
		handler.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		handler.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	router.HandleRequests()
}
