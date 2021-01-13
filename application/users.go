package application

import (
    "github.com/bismarkanes/web-go/domain"
    "github.com/bismarkanes/web-go/domain/repository"
)

type Users interface {
    GetUser(ID int) *domain.Users
    GetAllUser() []domain.Users
    UpdateUser(id int, name, email string) int64
    CreateUser(name, email string) *domain.Users
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

func (us *users) UpdateUser(id int, name, email string) int64 {
    val, err := us.usersRepo.Update(id, name, email)
    if err != nil {
        return 0
    }
    return val
}

func (us *users) CreateUser(name, email string) *domain.Users {
    user, err := us.usersRepo.Create(name, email)
    if err != nil {
        return nil
    }

    return &user
}

func NewUsers(ur repository.Users) Users {
    return &users{
        usersRepo: ur,
    }
}
