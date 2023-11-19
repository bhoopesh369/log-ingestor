package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint    `gorm:"primaryKey;autoIncrement;not null"`
	Name         string  `gorm:"size:255;not null"`
	Email        string  `gorm:"size:255;unique;not null"`
}
