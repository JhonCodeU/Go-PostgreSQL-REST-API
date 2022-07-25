package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"`
	Phone     string `gorm:"not null"`
	Tasks     []Task 
	
}