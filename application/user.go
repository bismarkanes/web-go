package application

import (
    "github.com/bismarkanes/web-go/domain/model"
    "github.com/bismarkanes/web-go/domain/repository"
)

type User interface {
    GetUser(ID int) *model.User
    GetAllUser() []model.User
    UpdateUser(id int, name, email string) int64
    CreateUser(name, email string) *model.User
    DeleteUser(id int) error
}

type user struct {
    userRepo repository.User
}

func (us *user) GetUser(ID int) *model.User {
    user, err := us.userRepo.Get(ID)
    if err != nil {
        return nil
    }

    return &user
}

func (us *user) GetAllUser() []model.User {
    user, err := us.userRepo.GetAll()
    if err != nil {
        return nil
    }

    return user
}

func (us *user) UpdateUser(id int, name, email string) int64 {
    val, err := us.userRepo.Update(id, name, email)
    if err != nil {
        return 0
    }
    return val
}

func (us *user) CreateUser(name, email string) *model.User {
    user, err := us.userRepo.Create(name, email)
    if err != nil {
        return nil
    }

    return &user
}

func (us *user) DeleteUser(id int) error {
    return us.userRepo.Delete(id)
}

func NewUser(ur repository.User) User {
    return &user{
        userRepo: ur,
    }
}
