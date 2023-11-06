package handlers

import (
	"github.com/DelaRicch/klock-ecommerce/server/database"
	"github.com/DelaRicch/klock-ecommerce/server/models"
	"github.com/gofiber/fiber/v2"
)

func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("This is Klock E-commerce web app")
}

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var existingUser models.User
	
    result := database.DB.Where("email = ?", user.Email).First(&existingUser)
    if result.RowsAffected > 0 {
        return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
            "message": "Email already exists",
        })
    }
	
	database.DB.Create(&user)
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func DeleteAllUsers(ctx *fiber.Ctx) error {
	if err := database.DB.Exec("DELETE FROM users WHERE role = 'USER' ").Error; err != nil {
        return err
    }

    return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "All users deleted",
	})
}