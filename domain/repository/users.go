package repository

import "github.com/bismarkanes/web-go/domain/model"

// Users .
type Users interface {
  Get(id int) (model.Users, error)
  GetAll() ([]model.Users, error)
}
