package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name       string `json:"name" gorm:"text;not null;default:null"`
	Email      string `json:"email" gorm:"text;not null;default:null"`
	Password   string `json:"password" gorm:"text;default:''"`
	UserID     string `json:"userId" gorm:"text;not null;default:null"`
	Role       string `json:"role" gorm:"text;not null;default:null"`
	RememberMe string `json:"remember" gorm:"bool;default:false"`
	Photo      string `json:"photo" gorm:"text;default:null"`
	Phone      string `json:"phone" gorm:"text;default:null"`
	SocialId   string `json:"socialId" gorm:"text;default:null"`
	Provider   string `json:"provider" gorm:"text;default:null"`
}

type ProductGalleryImage struct {
	gorm.Model
	ImageURL  string `json:"imageUrl" gorm:"not null;default:null"`
	ProductID string `json:"productId" gorm:"text;not null"`
}

type Product struct {
	gorm.Model

	ProductName               string                `json:"productName" gorm:"text;not null;default:null"`
	ProductDescription        string                `json:"productDescription" gorm:"text;not null;default:null"`
	ProductCategory           string                `json:"productCategory" gorm:"text;not null;default:null"`
	ProductPrice              string                `json:"productPrice" gorm:"uint;not null;default:0"`
	ProductDiscountPercentage string                `json:"productDiscountPercentage" gorm:"uint;not null;default:0"`
	ProductQuantity           string                `json:"productQuantity" gorm:"uint;not null;default:0"`
	ProductBrandName          string                `json:"productBrandName" gorm:"text;not null;default:null"`
	ProductCoverImage         string                `json:"productCoverImage" gorm:"not null;default:null"`
	ProductGalleryImages      []ProductGalleryImage `json:"productGalleryImages" gorm:"foreignKey:ProductID"`
	ProductID                 string                `json:"productId" gorm:"text;not null;default:null"`
}
