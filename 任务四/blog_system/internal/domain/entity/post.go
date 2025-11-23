package entity

import (
	"time"

	"gorm.io/gorm"
)
type Post struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:200;not null"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
func (p *Post) BeforeCreate(tx *gorm.DB) error {
    p.CreatedAt = time.Now()
    p.UpdatedAt = time.Now()
    return nil
}

func (p *Post) BeforeUpdate(tx *gorm.DB) error {
    p.UpdatedAt = time.Now()
    return nil
}