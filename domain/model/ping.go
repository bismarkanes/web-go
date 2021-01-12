package domain

import (
  "github.com/jinzhu/gorm"
)

// SalesTeam .
type Ping struct {
  gorm.Model
  Name string `gorm:"column:name"`
}
