package repository

import (
    "github.com/bismarkanes/web-go/domain/model"
)

// User .
type User interface {
    Get(id int) (model.User, error)
    GetAll() ([]model.User, error)
    Create(name, email string) (model.User, error)
    Update(id int, name, email string) (int64, error)
    Delete(id int) error
}
