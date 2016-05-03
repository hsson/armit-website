package app

import (
  "net/http"
  "fmt"
	"strings"
  "log"
  "io/ioutil"

  "github.com/dgrijalva/jwt-go"
)

// The secret used to sign the JWT tokens
// TODO: Replace the implementation of one single token to use multiple ones
var appSecretKey string

func InitAuth() {
  key, err := ioutil.ReadFile("app-secret.pwd")
  if err != nil {
    log.Fatalf("App authorization could not be initialized: %v", err)
  }
  appSecretKey = string(key)
  log.Printf("Authorization successfully initialized with key: %s", appSecretKey)
}

func Authenticator(inner http.Handler, authed bool) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !authed || checkAuth(w,r) {
      inner.ServeHTTP(w,r)
      return
    }
    w.WriteHeader(401)
    w.Write([]byte("401 Unauthorized\n"))
  })
}

// Checks if the request is authorized
func checkAuth(w http.ResponseWriter, r *http.Request) bool {
  // Header looks like: "Authorization: Bearer <token>"
  s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
    log.Println("Invalid Authorization header length used")
    return false
  }

  appToken := string(s[1])

  // Parse the token and check if it is valid. Very important to check for
  // the signing method
  token, err := jwt.Parse(appToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            log.Printf("Unauthorized signing method used: %v", token.Header["alg"])
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(appSecretKey), nil
    })

  if err == nil && token.Valid {
    log.Println("Successfully authenticated")
    return true
  } else {
    log.Printf("Bad authentication: %v", err)
    return false
  }
}
