package persistence

import (
    "github.com/bismarkanes/web-go/domain"
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

func (us *user) Get(id int) (domain.User, error) {
    user := domain.User{}

    err := us.dbConn.First(&user, id).Error

    return user, err
}

func (us *user) GetAll() ([]domain.User, error) {
    user := []domain.User{}

    err := us.dbConn.Find(&user).Error

    return user, err
}

func (us *user) Update(id int, name, email string) (int64, error) {
    // var val int
    user := domain.User{}

    result := us.dbConn.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "email": email})

    return result.RowsAffected, result.Error
}

func (us *user) Create(name, email string) (domain.User, error) {
    user := domain.User{
        Name:  name,
        Email: email,
    }

    result := us.dbConn.Create(&user)

    err := result.Scan(&user).Error
    return user, err
}
