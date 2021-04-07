package posts

import (
  "fmt"
  "net/http"
  "gorm.io/gorm"

  "post-api/db"
)

func Index(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
  if r.Method == "GET" {
    fmt.Fprintf(w, "hello")
  }
}

func Create(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
  if r.Method == "POST" {
    post := db.Post{Content: "Hello World"}
    Db.Create(&post)
  }
}
