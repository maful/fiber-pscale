package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/maful/fiber-pscale/models"
)

// TODO: Add validation
type createUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Website string `json:"website"`
}

type updateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Website string `json:"website"`
}

func CreateUser(c *fiber.Ctx) error {
	req := &createUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		Name:    req.Name,
		Email:   req.Email,
		Website: req.Website,
	}
	models.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"user": user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	models.DB.Find(&users)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"users": users,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"user": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// first, check if the user is exist
	user := models.User{}
	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	// second, parse the request body
	request := &updateUserRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// third, update the user
	updateUser := models.User{
		Name:    request.Name,
		Email:   request.Email,
		Website: request.Website,
	}
	models.DB.Model(&user).Updates(&updateUser)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"user": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	// first, check if the user is exist
	user := models.User{}
	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	// second, delete the user
	models.DB.Delete(&user)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
	})
}
