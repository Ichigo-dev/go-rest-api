package main

import (
  "post-api/db"
  "post-api/config"
  "net/http"
  _"fmt"
  _"reflect"
)

func main() {
  Db,_ := db.Db()

  server := http.Server{
    Addr: "127.0.0.1:8080",
    Handler: config.Route(Db),
  }
  server.ListenAndServe()
}
