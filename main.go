package main

import (
  "net/http"
  "log"

  "github.com/hsson/armit-website/app"
)

func main() {
  app.InitAuth()
  r := app.NewRouter()
  log.Println("Server starting...")
  log.Fatal(http.ListenAndServe(":5000", r))
}
