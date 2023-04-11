package main

import (
  "log"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql" 
  "fmt"
)

type Todo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Description bool `json:"completed"`
}

func main() {
  log.Println("hello")
  db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+ "3306" +")/"+os.Getenv("DB_NAME"))
	if err != nil {
    log.Fatal(err)
	}
	defer db.Close()

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("DB connected")
  }
}
