package db

import (
  "time"
)

type Post struct {
  Id uint `gorm:"primaryKey"`
  Content string `gorm:not null`
  CreatedAt time.Time
  UpdatedAt time.Time
}
