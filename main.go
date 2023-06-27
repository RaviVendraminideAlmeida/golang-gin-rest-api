package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

var todos = []todo{
	{ID: "1", Title: "Implement GET", Content: "Get all the elements of the list", Done: false},
	{ID: "2", Title: "Implement POST", Content: "Add an element to the list", Done: false},
	{ID: "3", Title: "Implement PUT", Content: "Update an element from the list", Done: false},
	{ID: "5", Title: "Implement PATCH", Content: "Mark an element from the as done", Done: false},
	{ID: "5", Title: "Implement Delete", Content: "Delete an element of the list", Done: false},
}

func remove(t []todo, i int) []todo {
	t[i] = t[len(t)-1]
	return t[:len(t)-1]
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": " not found"})
}

func markAsDone(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todos[i].Done = true
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func postTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func putTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	for _, t := range todos {
		if t.ID == newTodo.ID {
			t = newTodo
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todos = remove(todos, i)
			c.IndentedJSON(http.StatusOK, todos)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodo)
	router.PATCH("/todos/:id", markAsDone)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run("localhost:8080")
}
