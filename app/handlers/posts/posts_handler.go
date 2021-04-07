package posts

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"

  "post-api/db"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "hello")
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  post := db.Post{Content: "Hello World"}
  db.Db.Create(&post)
}
