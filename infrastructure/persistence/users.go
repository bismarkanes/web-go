package persistence

import (
  "github.com/bismarkanes/web-go/domain/model"
  "github.com/bismarkanes/web-go/domain/repository"
  "gorm.io/gorm"
)

// Users .
type users struct {
  dbConn *gorm.DB
}

// NewUsersRepo ...
func NewUsersRepo(dbConn *gorm.DB) repository.Users {
  return &users{
    dbConn: dbConn,
  }
}

func (us *users) Get(id int) (model.Users, error) {
  user := model.Users{}

  err := us.dbConn.First(&user, id).Error

  return user, err
}

func (us *users) GetAll() ([]model.Users, error) {
  users := []model.Users{}

  err := us.dbConn.Find(&users).Error

  return users, err
}
