package controllers

import (
	"fmt"
	"mongodb-go/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	users, err := models.Find()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    users,
	})
}

func CreateData(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err := models.Insert(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"Message": "success",
			"ID":      user.ID.Hex(), // Mengembalikan ID dalam format string hexadecimal
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not allowed")
	}
}

func UpdateData(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPut {
		idParams := c.Params("id")

		var updateUser models.User
		if err := c.BodyParser(&updateUser); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err := models.Update(idParams, &updateUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"Message": fmt.Sprintf("update successfully with id %s", idParams),
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not allowed")
	}
}

func DeleteData(c *fiber.Ctx) error {
	idParams := c.Params("id")
	models.Delete(idParams)

	return c.JSON(fiber.Map{
		"Message": fmt.Sprintf("Deleted Success with ID %s", idParams),
	})
}
