package application

import (
    "github.com/bismarkanes/web-go/domain"
    "github.com/bismarkanes/web-go/domain/repository"
)

type Users interface {
    GetUser(ID int) *domain.Users
    GetAllUser() []domain.Users
}

type users struct {
    usersRepo repository.Users
}

func (us *users) GetUser(ID int) *domain.Users {
    users, err := us.usersRepo.Get(ID)
    if err != nil {
        return nil
    }

    return &users
}

func (us *users) GetAllUser() []domain.Users {
    users, err := us.usersRepo.GetAll()
    if err != nil {
        return nil
    }

    return users
}

func NewUsers(ur repository.Users) Users {
    return &users{
        usersRepo: ur,
    }
}
