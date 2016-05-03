package app

import (
  "net/http"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  Authed      bool
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    false,
    Index,
  },
  Route{
    "Secured",
    "GET",
    "/secured",
    true,
    SecuredPage,
  },
}
