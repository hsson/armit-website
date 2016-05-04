package main

import (
  "net/http"
  "log"

  "github.com/hsson/armit-website/app"
  "github.com/hsson/armit-website/app/handlers"
)

const ServerPort = ":8080"

func main() {
  app.InitAuth()
  handlers.InitPartners() // TODO: Replace partners json with database
  r := app.NewRouter()
  log.Printf("Server starting on localhost%s...", ServerPort)
  log.Fatal(http.ListenAndServe(ServerPort, r))
}
