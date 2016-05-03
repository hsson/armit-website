package main

import (
  "net/http"
  "log"

  "github.com/hsson/armit-website/app"
)

const ServerPort = ":8080"

func main() {
  app.InitAuth()
  r := app.NewRouter()
  log.Printf("Server starting on localhost%s...", ServerPort)
  log.Fatal(http.ListenAndServe(ServerPort, r))
}
