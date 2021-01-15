package model

// User .
type User struct {
    ModelBase
    Name  string `gorm:"column:name" json:"name"`
    Email string `gorm:"column:email" json:"email"`
}
