package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: 1, Item: "Get up 7 am", Completed: false},
	{Id: 2, Item: "Morning Workout", Completed: true},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context){
	var newTodo todo

	if err := c.BindJSON(&newTodo) ; err != nil{
		return 
	}
	todos  = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	fmt.Println(todos[1].Item)
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.POST("/todo", addTodo)
	r.Run("localhost:8080")

}
