package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
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
	{ID: "4", Title: "Implement PATCH", Content: "Mark an element from the list as done", Done: false},
	{ID: "5", Title: "Implement Delete", Content: "Delete an element of the list", Done: false},
}

func remove(t []todo, i int) []todo {
	t[i] = t[len(t)-1]
	return t[:len(t)-1]
}

func HomePageHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H {
		"message" : "Welcome to my TODO list App! :)",
	})
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
}

func MarkAsDone(c *gin.Context) {
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

func PostTodo(c *gin.Context) {
	var newTodo todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newTodo.ID = xid.New().String()
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func PutTodo(c *gin.Context) {
	id := c.Param("id")
	var newTodo todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			index = 1
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "TODO not found",
		})
		return
	}

	index++

	todos[index].Content = newTodo.Content
	todos[index].Done = newTodo.Done
	todos[index].Title = newTodo.Title
	c.JSON(http.StatusOK, todos[index])
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todos = remove(todos, i)
			c.JSON(http.StatusOK, gin.H{
				"message": "TODO deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "TODO not found"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomePageHandler)
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodoByID)
	router.PUT("/todos/:id", PutTodo)
	router.POST("/todos", PostTodo)
	router.PATCH("/todos/:id", MarkAsDone)
	router.DELETE("/todos/:id", DeleteTodo)
	router.Run("localhost:8080")
}
