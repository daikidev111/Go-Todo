package main

import (
  "log"
  // "os"
  "github.com/Go-Todo/Config"
  "github.com/Go-Todo/Routes"
  "database/sql"
)

func main() {
  db, err := sql.Open("mysql", Config.DBUri(Config.BuildDBConfig()))
	if err != nil {
    log.Fatal(err) // print the log and terminates the application with os.Exit(1)
	}

  Config.DB = db
	defer Config.DB.Close()

  r := Routes.SetupRouter()
  r.Run()
}
