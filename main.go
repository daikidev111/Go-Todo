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
  db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+ "3306" +")/"+os.Getenv("DB_NAME"))
	if err != nil {
    log.Fatal(err)
	}
	defer db.Close()

  // test the DB connection
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("DB connected -> pong")
  }

}
