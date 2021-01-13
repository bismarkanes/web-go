package domain

// Users .
type Users struct {
    ModelCommon
    Name  string `gorm:"column:name" json:"name"`
    Email string `gorm:"column:email" json:"email"`
}
