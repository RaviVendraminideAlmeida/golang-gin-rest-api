package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Test_GetTodos(t *testing.T){
	expected := []todo{
		{ID: "1", Title: "Implement GET", Content: "Get all the elements of the list", Done: false},
		{ID: "2", Title: "Implement POST", Content: "Add an element to the list", Done: false},
		{ID: "3", Title: "Implement PUT", Content: "Update an element from the list", Done: false},
		{ID: "5", Title: "Implement PATCH", Content: "Mark an element from the list as done", Done: false},
		{ID: "5", Title: "Implement Delete", Content: "Delete an element of the list", Done: false},
	}
	got := GetTodos()

	if (expected != got){
		t.Errorf("Expected: ")

	}
}

func Test_getTodoById(t *testing.T){
	return
}

func Test_getTodoByID(t *testing.T){
	return
}

func Test_putTodo(t *testing.T){
	return
}

func Test_postTodo(t *testing.T){
	return
}

func Test_deleteTodo(t *testing.T){
	return
}

func Test_markAsDone(t *testing.T){
	return
}




func get