package app

import (
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

// Creates a router that handles the routes specified in the routes.go file
func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  apiRouter := router.PathPrefix("/api").Subrouter()
  for _, route := range routes {
    var handler http.Handler

    // Run every request through an authenticator
    handler = Authenticator(Logger(route.HandlerFunc, route.Name), route.Authed)

    apiRouter.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
  log.Printf("Backend router successfully initialized with %d routes.", len(routes))

  fileServer := Logger(http.FileServer(http.Dir("static")), "static")
  router.PathPrefix("/").Handler(fileServer)
  log.Println("Web router successfully initialized.")
  return router
}
