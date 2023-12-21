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
	Location      string `json:"location" gorm:"text;default:null"`
	SocialId   string `json:"socialId" gorm:"text;default:null"`
	Provider   string `json:"provider" gorm:"text;default:null"`
}

type UserProfile struct {
	Name string
	Email string
	UserID string
	Role string
	Photo string
	Phone string
	Location string
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
	ProductPrice              float64               `json:"productPrice" gorm:"float64;not null;default:0"`
	ProductDiscountPercentage float64               `json:"productDiscountPercentage" gorm:"float64;not null;default:0"`
	ProductQuantity           uint                  `json:"productQuantity" gorm:"uint;not null;default:0"`
	ProductBrandName          string                `json:"productBrandName" gorm:"text;not null;default:null"`
	ProductCoverImage         string                `json:"productCoverImage" gorm:"not null;default:null"`
	ProductGalleryImages      []ProductGalleryImage `json:"productGalleryImages" gorm:"foreignKey:ProductID"`
	ProductID                 string                `json:"productId" gorm:"text;not null;default:null"`
	ProductsRemaining         uint                  `json:"productsRemaining" gorm:"uint;not null;default:0"`
	ProductsSold              uint                  `json:"productsSold" gorm:"uint;not null;default:0"`
}

type ProductForFrontend struct {
	ProductBrandName          string
	ProductCategory           string
	ProductCoverImage         string
	ProductDescription        string
	ProductDiscountPercentage float64
	ProductGalleryImages      []ProductGalleryImage
	ProductID                 string
	ProductName               string
	ProductPrice              float64
	ProductQuantity           uint
	ProductsRemaining         uint
	ProductsSold              uint
}
