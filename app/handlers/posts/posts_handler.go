package posts

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "github.com/antonholmquist/jason"

  "post-api/db"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "hello")
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  v, _ := jason.NewObjectFromReader(r.Body)
  content, _ := v.GetString("content")

  post := db.Post{Content: content}
  db.Db.Create(&post)
}
