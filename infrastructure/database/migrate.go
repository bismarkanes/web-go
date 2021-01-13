package database

import (
    "github.com/bismarkanes/web-go/domain/model"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    Users := model.Users{}

    return db.AutoMigrate(&Users)
}
