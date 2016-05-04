package handlers

import (
  "encoding/json"
  "log"
  "strconv"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/hsson/armit-website/app/models"
  "github.com/gorilla/mux"
)

var partners models.Partners

// TODO: Replace with database
func InitPartners() {
  partnerJson, err := ioutil.ReadFile("data/partners.json")
  if err != nil {
    log.Fatal("Partners could not be initialized.")
  }
  err = json.Unmarshal(partnerJson, &partners)
  if err != nil {
    log.Fatal("Partners could not be initialized.")
  }
}

// GET: /api/partners
func Partners(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(partners); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
  }
}

// GET: /api/partners/id:1
func Partner(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
  }

  if id + 1 > len(partners) {
    log.Println("Partner index out of bounds")
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "400 - Index does not exist")
    return
  }
  if err := json.NewEncoder(w).Encode(partners[id]); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
  }
}
