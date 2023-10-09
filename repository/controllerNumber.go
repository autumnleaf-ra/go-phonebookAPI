package repository

import (
	"net/http"
	"strconv"

	"github.com/autumnleaf-ra/phonebook-api/models"
	"github.com/gofiber/fiber/v2"
)

/* Number */

func (r *Repository) GetNumber(c *fiber.Ctx) error {
	var getData []models.PhoneNumber

	/* Query Get Data */
	if err := r.DB.Find(&getData).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Error getting data !",
			},
		)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": getData,
	})
}

func (r *Repository) AddNumber(c *fiber.Ctx) error {
	var phone models.PhoneNumber
	err := c.BodyParser(&phone)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "RequestFaied"},
		)
		return err
	}

	/* Check Duplicate Number */
	if err := r.DB.Where("number = ?", phone.Number).First(&phone).Error; err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Phone number already exists",
		})
	}

	/* Query Add Number */
	if err := r.DB.Create(&phone).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Phone number has been add",
	})

	return nil
}

func (r *Repository) DeleteNumber(c *fiber.Ctx) error {
	/* Getting Params ID */
	id, _ := strconv.Atoi(c.Params("id"))
	phone := models.PhoneNumber{
		Id: uint(id),
	}

	/* Query Delete Number */
	delete := r.DB.Delete(&phone)
	if delete.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Phone number not found !",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Phone number deleted successfully",
	})
}

func (r *Repository) DetailNumber(c *fiber.Ctx) error {
	/* Getting Params ID */
	id, _ := strconv.Atoi(c.Params("id"))
	phone := models.PhoneNumber{
		Id: uint(id),
	}

	/* Query Detail Number */
	if err := r.DB.First(&phone).Error; err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Phone number not found !",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": phone,
	})

}
