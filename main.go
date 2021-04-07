package main

import (
  "post-api/db"
  "post-api/app/handlers/posts"

  "net/http"
  "github.com/julienschmidt/httprouter"
  _"fmt"
)

func main() {
  Db,_ := db.Db()

  route := httprouter.New()
  route.GET("/posts", func(w http.ResponseWriter, h *http.Request, _ httprouter.Params){ posts.Index(w, h, Db)})
  route.POST("/posts", func(w http.ResponseWriter, h *http.Request, _ httprouter.Params){ posts.Create(w, h, Db)})

  server := http.Server{
    Addr: "127.0.0.1:8080",
    Handler: route,
  }
  server.ListenAndServe()
}
