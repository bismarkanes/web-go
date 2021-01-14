package repository

import "github.com/bismarkanes/web-go/domain"

// User .
type User interface {
    Get(id int) (domain.User, error)
    GetAll() ([]domain.User, error)
    Create(name, email string) (domain.User, error)
    Update(id int, name, email string) (int64, error)
    Delete(id int) error
}
