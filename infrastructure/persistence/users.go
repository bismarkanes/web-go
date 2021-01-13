package persistence

import (
    "github.com/jinzhu/gorm"
)

// Users .
type Users struct {
    dbConn *gorm.DB
}

// NewOrder ...
func NewUsersRepo(dbConn *gorm.DB) repo {
    return &Users{
        dbConn: dbConn,
    }
}
