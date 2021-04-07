package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var Db *gorm.DB

func init() {
  var err error
  dsn := "host=localhost user=postgres dbname=zero_book_demo sslmode=disable"
  Db, err = gorm.Open(postgres.Open(dsn))
  if err != nil {
    panic(err)
  }
  Db.AutoMigrate(Post{})
}
