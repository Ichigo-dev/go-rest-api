package main

import (
  "post-api/db"
  "post-api/app/handlers/posts"

  "net/http"
)

func main() {
  Db,_ := db.Db()
  Db.AutoMigrate(&db.Post{})

  server := http.Server{
    Addr: "127.0.0.1:8080",
  }

  http.HandleFunc("/posts", posts.Index)

  server.ListenAndServe()
}
