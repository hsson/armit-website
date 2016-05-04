package app

import (
  "net/http"

  "github.com/hsson/armit-website/app/handlers"
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
    handlers.Index,
  },
  Route{
    "Secured",
    "GET",
    "/secured",
    true,
    handlers.SecuredPage,
  },
  Route{
    "Partners",
    "GET",
    "/partners",
    false,
    handlers.Partners,
  },
  Route{
    "Partner",
    "GET",
    "/partners/{id:[0-9]+}",
    false,
    handlers.Partner,
  },
}
