package posts

import (
  _"fmt"
  "net/http"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/antonholmquist/jason"

  "post-api/db"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var posts []db.Post
  db.Db.Find(&posts)
  json.NewEncoder(w).Encode(posts)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  v, _ := jason.NewObjectFromReader(r.Body)
  content, _ := v.GetString("content")

  post := db.Post{Content: content}
  db.Db.Create(&post)
  json.NewEncoder(w).Encode(&post)
}
