package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/todos", CommonHeaders(TodoIndex, "todo-index"))
	router.GET("/todos/:todoId", IDShouldBeInt(TodoShow, "todo-show"))
	router.POST("/todos", CommonHeaders(TodoCreate, "todo-create"))
	router.DELETE("/todos/:todoId", IDShouldBeInt(TodoDelete, "todo-delete"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
