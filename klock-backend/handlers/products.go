package handlers

import (
	"fmt"

	"github.com/DelaRicch/klock-ecommerce/klock-backend/database"
	"github.com/DelaRicch/klock-ecommerce/klock-backend/lib"
	"github.com/DelaRicch/klock-ecommerce/klock-backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddNewProduct(ctx *fiber.Ctx) error {
	newProduct := new(models.Product)
	req, _ := ctx.MultipartForm()

	if err := ctx.BodyParser(newProduct); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
			"success": false,
		})
	}

	// Validate input using the validator library
	validate := validator.New()
	if err := validate.Struct(newProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}


	// Check for existing product in the system
	var existingProduct models.Product
	res := database.DB.Where("product_name = ?", newProduct.ProductName).First(&existingProduct)
	if res.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": fmt.Sprintf("%v already exist", newProduct.ProductName),
			"Success": false,
		})
	}
	// Generate Product ID
	newProduct.ProductID = lib.GenerateID("KLOCK-PRODUCT")

	newProduct.ProductsRemaining = newProduct.ProductQuantity - newProduct.ProductsSold

	// Upload cover image to Cloudinary
	coverImage := req.File["productCoverImage"][0]
	coverImageUrl, err := lib.UploadToCloudinary(coverImage, newProduct.ProductID, "cover-image")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error uploading cover image",
			"success": false,
		})
	}

	// Upload gallery images to Cloudinary
	var galleryImageURLs []string
	for _, galleryImage := range req.File["productGalleryImages"] {
		imageUrl, err := lib.UploadToCloudinary(galleryImage, newProduct.ProductID, "gallery-images")
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error uploading gallery image",
				"success": false,
			})
		}
		galleryImageURLs = append(galleryImageURLs, imageUrl)
	}

	// Append cover image to newProduct
	newProduct.ProductCoverImage = coverImageUrl

	// Append gallery images to newProduct
	newProduct.ProductGalleryImages = make([]models.ProductGalleryImage, 0)
	for _, imageUrl := range galleryImageURLs {
		galleryImage := models.ProductGalleryImage{
			ImageURL: imageUrl,
		}
		newProduct.ProductGalleryImages = append(newProduct.ProductGalleryImages, galleryImage)
	}

	// Add new product to database
	result := database.DB.Create(&newProduct)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error adding new product to database",
			"success": false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          fmt.Sprintf("Successfully added %s", newProduct.ProductName),
		"success":          true,

	})
}

func ExcludeFields(product models.Product) models.ProductForFrontend {
	return models.ProductForFrontend{
		ProductBrandName:       product.ProductBrandName,
		ProductCategory:        product.ProductCategory,
		ProductCoverImage:      product.ProductCoverImage,
		ProductDescription:     product.ProductDescription,
		ProductDiscountPercentage: product.ProductDiscountPercentage,
		ProductGalleryImages:   product.ProductGalleryImages,
		ProductID:              product.ProductID,
		ProductName:            product.ProductName,
		ProductPrice:           product.ProductPrice,
		ProductQuantity:        product.ProductQuantity,
		ProductsRemaining:      product.ProductsRemaining,
		ProductsSold:           product.ProductsSold,
	}
}


func ListProducts(ctx *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Preload("ProductGalleryImages").Find(&products)

	// Create a new slice with excluded fields for the frontend
	productsForFrontend := make([]models.ProductForFrontend, len(products))
	for i, product := range products {
		productsForFrontend[i] = ExcludeFields(product)
	}

	return ctx.Status(fiber.StatusOK).JSON(productsForFrontend)
}

func DeleteAllProducts(ctx *fiber.Ctx) error {
    tx := database.DB.Begin()

    // Delete associated gallery images
    if err := tx.Exec("DELETE FROM product_gallery_images").Error; err != nil {
        tx.Rollback()
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Delete products
    if err := tx.Exec("DELETE FROM products").Error; err != nil {
        tx.Rollback()
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    tx.Commit()

    return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
        "message": "All products and associated gallery images deleted",
        "success": true,
    })
}

