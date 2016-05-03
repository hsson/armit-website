package app

import (
  "net/http"
  "log"
  "os"
  "github.com/gorilla/mux"
)

const StaticFolder = "static"

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

  fileServer := Logger(noDirListing(http.FileServer(http.Dir(StaticFolder))), StaticFolder)
  router.PathPrefix("/").Handler(fileServer)
  log.Println("Web router successfully initialized.")
  return router
}

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  	if f, err := os.Stat(StaticFolder + r.URL.Path); r.URL.Path != "/" && (err != nil || f.IsDir()) {
  		http.NotFound(w, r)
  		return
  	}
		h.ServeHTTP(w, r)
	})
}
