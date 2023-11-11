package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	
	Name string `json:"name" gorm:"text;not null;default:null"`
	Email string `json:"email" gorm:"text;not null;default:null"`
	Password string `json:"password" gorm:"text;default:''"`
	UserID string `json:"userId" gorm:"text;not null;default:null"`
	Role string `json:"role" gorm:"text;not null;default:null"`
	RememberMe string `json:"remember" gorm:"bool;default:false"`
	Photo string `json:"photo" gorm:"text;default:null"`
	Phone string `json:"phone" gorm:"text;default:null"`
	SocialId string `json:"socialId" gorm:"text;default:null"`
	Provider string `json:"provider" gorm:"text;default:null"`
}
