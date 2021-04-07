package main

import (
  "post-api/db"
)

func main() {
  Db,_ := db.Db()
  Db.AutoMigrate(&db.Post{})
}
