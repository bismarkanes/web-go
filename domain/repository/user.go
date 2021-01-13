package repository

import "github.com/bismarkanes/web-go/domain"

// Users .
type Users interface {
    Get(id int) (domain.Users, error)
    GetAll() ([]domain.Users, error)
}
