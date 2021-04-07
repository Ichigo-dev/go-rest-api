package config

import (
  "github.com/julienschmidt/httprouter"
  "gorm.io/gorm"
  "net/http"
  "post-api/app/handlers/posts"
  "fmt"
)

var route *httprouter.Router

func init() {
  fmt.Println("hello")
}

func Route(Db *gorm.DB) (route *httprouter.Router) {
  route = httprouter.New()
  route.GET("/posts", func(w http.ResponseWriter, h *http.Request, _ httprouter.Params){ posts.Index(w, h, Db)})
  route.POST("/posts", func(w http.ResponseWriter, h *http.Request, _ httprouter.Params){ posts.Create(w, h, Db)})
  return
}
