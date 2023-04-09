package controllers

import (
	"net/http"
	"strconv"
	"github.com/Go-Todo/Models"
	"github.com/gin-gonic/gin"
  "log"
)

// requestTodo represents the request body for creating or updating a Todo.
type requestTodo struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
}

// responseTodo represents the response body for a Todo.
type responseTodo struct {
  ID          uint   `json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
}

// GetTodos is a request handler for getting all Todos.
func GetTodos(c *gin.Context) {
  // invoke GetAllTodos() from models.go to retrieve all todos
	todos, err := models.GetAllTodos()
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return
	}
  
  // convert the todos to responseTodo type
  res := []responseTodo{}

  // iterate through the todos and append the responseTodo type to res
  for _, todos := range todos {
    res = append(res, responseTodo{
      ID: todos.ID,
      Title: todos.Title,
      Description: todos.Description,
    })
  }
  
  // return the responseTodo type along with the status code
	c.JSON(http.StatusOK, res)
}

// CreateTodo is a request handler for creating a new Todo.
func CreateTodo(c *gin.Context) {
    var req requestTodo
    if err := c.ShouldBindJSON(&req); err != nil {
      log.Println(err)
      c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
      return
    } 

    err := models.CreateTodo(req.Title, req.Description)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a todo task"})
        return
    }
    res := responseTodo{
      Title: req.Title,
      Description: req.Description,
    }

    c.JSON(http.StatusCreated, res)
}

// GetTodoByID is a request handler for getting a Todo by its ID.
func GetTodoByID(c *gin.Context) {
  id, err:= strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  todo, err := models.GetTodoByID(uint(id))
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return
	}

  res := responseTodo{
    ID: todo.ID,
    Title: todo.Title,
    Description: todo.Description,
  }

	c.JSON(http.StatusOK, res)
}

// UpdateTodo is a request handler for updating a Todo.
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  var req requestTodo
  if err := c.ShouldBindJSON(&req); err != nil {
    log.Println(err)
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    return
  }

  req = requestTodo{
    ID: uint(id),
    Title: req.Title,
    Description: req.Description,
  }

	err = models.UpdateTodo(req.ID, req.Title, req.Description)
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return
	}

  c.JSON(http.StatusOK, gin.H{"success":"Successfully Updated"})
}

// DeleteTodo is a request handler for deleting a Todo.
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

	err = models.DeleteTodo(uint(id))
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return 
	} 		
  c.JSON(http.StatusOK, gin.H{"success":"Successfully Deleted"})
}

