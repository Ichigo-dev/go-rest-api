package main

import (
  "zero-book-demo/db"
)

func main() {
  Db,_ := db.Db()
  Db.AutoMigrate(&db.Post{})
}
