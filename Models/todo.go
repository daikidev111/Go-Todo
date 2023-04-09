package Models

import (
	"fmt"

	"github.com/Go-Todo/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTodos() ([]*Todo, error) {
  todos := make([]*Todo, 0)

  rows, err := Config.DB.Query("SELECT * FROM todos")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    todo := &Todo{} // initialise the struct Todo
    err := rows.Scan(&todo.ID, &todo.Title, &todo.Description)
    if err != nil {
      return nil, err
    }
    todos = append(todos, todo)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  if len(todos) == 0 {
    fmt.Println("No todos found")
    return nil, nil
  }

  return todos, nil
}

func CreateTodo(title, description string) error {
    stmt, err := Config.DB.Prepare("INSERT INTO todos(title, description) VALUES(?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(title, description)
    if err != nil {
        return err
    }

    return nil
}

func GetTodoByID(id uint) (*Todo, error) {
    todo := &Todo{}

    row := Config.DB.QueryRow("SELECT * FROM todos WHERE id = ?", id)
    err := row.Scan(&todo.ID, &todo.Title, &todo.Description)
    if err != nil {
        return nil, err
    }

    return todo, nil
}

func UpdateTodo(id uint, title, description string) error {
    _, err := GetTodoByID(id)
    if err != nil {
      return err
    }

    stmt, err := Config.DB.Prepare("UPDATE todos SET title = ?, description = ? WHERE id = ?")
    if err != nil {
        return err
    }

    defer stmt.Close()

    if _, err = stmt.Exec(title, description, id); err != nil {
      return err 
    }

    return nil
}

func DeleteTodo(id uint) error {
    stmt, err := Config.DB.Prepare("DELETE FROM todos WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    return nil
}
