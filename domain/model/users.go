package model

import (
  "github.com/jinzhu/gorm"
)

// Users .
type Users struct {
  gorm.Model
  Name  string `gorm:"column:name"`
  Email string `gorm:"column:email"`
}
