package domain

import (
    "gorm.io/gorm"
    "time"
)

type ModelCommon struct {
    ID        uint           `gorm:"column:id" json:"id"`
    CreatedAt time.Time      `json:"createdAt"`
    UpdatedAt time.Time      `json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
