package database

import (
    "github.com/bismarkanes/web-go/domain"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    Users := domain.User{}

    return db.AutoMigrate(&Users)
}
