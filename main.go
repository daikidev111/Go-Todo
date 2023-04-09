package main

import (
  "log"
  // "os"
  "github.com/Go-Todo/Config"
  "github.com/Go-Todo/Routes"
  "database/sql"
)

func main() {
  db, err := sql.Open("mysql", config.DBUri(config.BuildDBConfig()))
	if err != nil {
    log.Fatal(err) // print the log and terminates the application with os.Exit(1)
	}

  config.DB = db
	defer config.DB.Close()

  r := routes.SetupRouter()
  r.Run()
}
