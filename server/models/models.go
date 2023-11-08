package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	
	Name string `json:"Name" gorm:"text;not null;default:null"`
	Email string `json:"Email" gorm:"text;not null;default:null"`
	Password string `json:"Password" gorm:"text;not null;default:null"`
	UserID string `json:"UserId" gorm:"text;not null;default:null"`
	Role string `json:"Role" gorm:"text;not null;default:null"`
}