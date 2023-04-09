package routes

import (
	"net/http"

	"github.com/Go-Todo/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()  // create a new gin router

	// define a GET route for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// define routes for Todos
	r.GET("/", controllers.GetTodos)         // get all Todos
	r.GET("/todo/:id", controllers.GetTodoByID)  // get a Todo by ID
	r.POST("/todo/create", controllers.CreateTodo)  // create a new Todo
	r.PUT("/todo/update/:id", controllers.UpdateTodo)  // update a Todo
	r.DELETE("/todo/delete/:id", controllers.DeleteTodo)  // delete a Todo

	// define fallback routes for when a route doesn't match
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})

	return r  // return the router
}
