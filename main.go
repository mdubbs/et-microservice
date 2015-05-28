package main

import (
  "log"
  "net/http"
  "github.com/mdubbs/et-microservice/db"
)

func main() {
  db.InitializeDb()
  
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}