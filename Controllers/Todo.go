package Controllers

import (
	"net/http"
	"strconv"
	"github.com/Go-Todo/Models"
	"github.com/gin-gonic/gin"
  "log"
)


type requestTodo struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
}

type responseTodo struct {
  ID          uint   `json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
}

func GetTodos(c *gin.Context) {
	todos, err := Models.GetAllTodos()
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return
	}

  res := []responseTodo{}
  for _, todos := range todos {
    res = append(res, responseTodo{
      ID: todos.ID,
      Title: todos.Title,
      Description: todos.Description,
    })
  }

	c.JSON(http.StatusOK, res)
}

func CreateTodo(c *gin.Context) {
    var req requestTodo
    if err := c.ShouldBindJSON(&req); err != nil {
      log.Println(err)
      c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
      return
    } 

    err := Models.CreateTodo(req.Title, req.Description)
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

func GetTodoByID(c *gin.Context) {
  id, err:= strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  todo, err := Models.GetTodoByID(uint(id))
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

	err = Models.UpdateTodo(req.ID, req.Title, req.Description)
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return
	}

  c.JSON(http.StatusOK, gin.H{"success":"Successfully Updated"})
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

	err = Models.DeleteTodo(uint(id))
	if err != nil {
    log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
    return 
	} 		
  c.JSON(http.StatusOK, gin.H{"success":"Successfully Deleted"})
}

