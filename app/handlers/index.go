package handlers

import (
  "fmt"
  "net/http"
)

var (
	version = `======================================
= ArmIT Backend: v0.0                =
=                                    =
= Alexander Håkansson                =
= alexander@hakansson.xyz            =
= github.com/hsson/armit-website     =
======================================`
)

// GET: /api/
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, version)
}

// GET: /api/secured
func SecuredPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "SECURED PAGE")
}
