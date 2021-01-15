package persistence

import (
    "github.com/bismarkanes/web-go/domain/model"
    "github.com/bismarkanes/web-go/domain/repository"
    "gorm.io/gorm"
)

// User .
type user struct {
    dbConn *gorm.DB
}

// NewUserRepo ...
func NewUserRepo(dbConn *gorm.DB) repository.User {
    return &user{
        dbConn: dbConn,
    }
}

func (us *user) Get(id int) (model.User, error) {
    user := model.User{}

    err := us.dbConn.First(&user, id).Error

    return user, err
}

func (us *user) GetAll() ([]model.User, error) {
    user := []model.User{}

    err := us.dbConn.Find(&user).Error

    return user, err
}

func (us *user) Update(id int, name, email string) (int64, error) {
    // var val int
    user := model.User{}

    result := us.dbConn.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "email": email})

    return result.RowsAffected, result.Error
}

func (us *user) Create(name, email string) (model.User, error) {
    user := model.User{
        Name:  name,
        Email: email,
    }

    result := us.dbConn.Create(&user)

    err := result.Scan(&user).Error
    return user, err
}

func (us *user) Delete(id int) error {
    user := model.User{}

    result := us.dbConn.Model(&user).Delete("id = ?", id)

    return result.Error
}
