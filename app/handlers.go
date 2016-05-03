package app

import (
  "fmt"
  "net/http"
)
// GET: /
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to armit backend!")
}

// GET: /secured
func SecuredPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "SECURED PAGE")
}
