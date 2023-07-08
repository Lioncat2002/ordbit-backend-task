package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	UserID         uint
	Name           string `gorm:"size:255;"`
	Description    string `gorm:"size:255;"`
	Tag            string `gorm:"size:255;"`
	CurrentOwnerID uint
}
