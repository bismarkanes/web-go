package domain

// User .
type User struct {
    ModelCommon
    Name  string `gorm:"column:name" json:"name"`
    Email string `gorm:"column:email" json:"email"`
}
