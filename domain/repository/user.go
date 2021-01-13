package repository

import "github.com/bismarkanes/web-go/domain"

// Users .
type Users interface {
    Get(id int) (domain.Users, error)
    GetAll() ([]domain.Users, error)
    Create(name, email string) (domain.Users, error)
    Update(id int, name, email string) (int64, error)
}
