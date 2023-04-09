package models

import (
	"log"

	"github.com/Go-Todo/Config"
	_ "github.com/go-sql-driver/mysql"
)

// GetAllTodos fetches all todos from the database
func GetAllTodos() ([]*Todo, error) {
  // initialise the slice of Todo
  todos := make([]*Todo, 0)
  
  // query the database
  rows, err := config.DB.Query("SELECT * FROM todos")
  if err != nil {
    return nil, err
  }
  // close the statement when the function returns or panics
  defer rows.Close()

  // iterate over the rows
  for rows.Next() {
    todo := &Todo{} // initialise the struct Todo
    // scan the row into the struct todo 
    err := rows.Scan(&todo.ID, &todo.Title, &todo.Description)
    if err != nil {
      return nil, err
    }
    // append the todo to the slice todos
    todos = append(todos, todo)
  }
  
  if err = rows.Err(); err != nil {
    return nil, err
  }
  
  // if empty then return nil
  if len(todos) == 0 {
    log.Println("No todos found")
    return nil, nil
  }

  return todos, nil
}

// CreateTodo creates a new todo in the database
func CreateTodo(title, description string) error {
    // prepare the statement
    stmt, err := config.DB.Prepare("INSERT INTO todos(title, description) VALUES(?, ?)")
    if err != nil {
        return err
    }
    // close the statement when the function returns or panics
    defer stmt.Close()
    
    // execute the statement
    _, err = stmt.Exec(title, description)
    if err != nil {
        return err
    }

    return nil
}

// GetTodoByID fetches a todo from the database by id
func GetTodoByID(id uint) (*Todo, error) {
    // initialise the struct Todo
    todo := &Todo{}
    
    // query the database
    row := config.DB.QueryRow("SELECT * FROM todos WHERE id = ?", id)
    
    // scan the row into the struct todo
    err := row.Scan(&todo.ID, &todo.Title, &todo.Description)
    if err != nil {
        return nil, err
    }

    return todo, nil
}

// UpdateTodo updates a todo in the database
func UpdateTodo(id uint, title, description string) error {
    _, err := GetTodoByID(id) // check if todo exists
    if err != nil {
      return err
    }
    
    // update todo
    stmt, err := config.DB.Prepare("UPDATE todos SET title = ?, description = ? WHERE id = ?")
    if err != nil {
        return err
    }
    
    // close the statement when the function returns or panics
    defer stmt.Close()
    
    // execute the statement
    if _, err = stmt.Exec(title, description, id); err != nil {
      return err 
    }

    return nil
}

// DeleteTodo deletes a todo from the database
func DeleteTodo(id uint) error {
    stmt, err := config.DB.Prepare("DELETE FROM todos WHERE id = ?") // prepare the statement
    if err != nil {
        return err
    }

    // close the statement when the function returns or panics
    defer stmt.Close()
    
    // execute the statement
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    return nil
}
