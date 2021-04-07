package main

import (
  "post-api/config"
  "net/http"
  _"fmt"
  _"reflect"
)

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
    Handler: config.Route,
  }
  server.ListenAndServe()
}
