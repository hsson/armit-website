package app

import (
  "log"
  "net/http"
  "time"

  "github.com/fatih/color"
)

func Logger(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    inner.ServeHTTP(w,r)

    log.Printf(
      color.GreenString(r.Method) + "\t" +
      color.RedString(r.RequestURI) + "\t%s\t%s",
      name,
      time.Since(start),
    )
  })
}
