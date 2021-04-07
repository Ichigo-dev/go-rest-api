package posts

import (
  "fmt"
  "net/http"
  "gorm.io/gorm"
  _"github.com/julienschmidt/httprouter"

  "post-api/db"
)

func Index(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
  fmt.Fprintf(w, "hello")
}

func Create(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
    post := db.Post{Content: "Hello World"}
    Db.Create(&post)
}
