package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Items     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Items: "utem 1", Completed: true},
	{ID: "2", Items: "utem 2", Completed: true},
	{ID: "3", Items: "utem 3", Completed: false},
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
