package Routes

import (
	"net/http"

	"github.com/Go-Todo/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.GET("/", Controllers.GetTodos)
  r.GET("/todo/:id", Controllers.GetTodoByID)
  r.POST("/todo/create", Controllers.CreateTodo)
  r.PUT("/update/:id", Controllers.UpdateTodo)
  r.DELETE("/delete/:id", Controllers.DeleteTodo)
  r.NoRoute(func(c *gin.Context) {
      c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
  })
  r.NoMethod(func(c *gin.Context) {
    c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
  })

	return r
}
