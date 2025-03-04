package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID        uint           `gorm:"primaryKey"`
    Username      string     `gorm:"not null"`
    Email     string         `gorm:"unique;not null"`
    Password  string         `gorm:"not null"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}