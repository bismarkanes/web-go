package model

import (
  "gorm.io/gorm"
)

// Users .
type Users struct {
  gorm.Model
  Name  string `gorm:"column:name" json:"name"`
  Email string `gorm:"column:email" json:"email"`
}
