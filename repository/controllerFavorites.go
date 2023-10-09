package repository

import (
	"net/http"
	"strconv"

	"github.com/autumnleaf-ra/phonebook-api/models"
	"github.com/gofiber/fiber/v2"
)

func (r *Repository) GetNumberFavorites(c *fiber.Ctx) error {
	var FavoriteNumber []models.FavoriteNumber

	/* Query */
	if err := r.DB.Preload("PhoneNumber").Find(&FavoriteNumber).Error; err != nil {
		c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": FavoriteNumber,
	})
}

func (r *Repository) AddFavorites(c *fiber.Ctx) error {
	phone := models.FavoriteNumber{}
	/* Get Param ID */
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	phone.PhoneID = uint(id)

	/* Check Duplicate Number */
	if err := r.DB.Preload("PhoneNumber").Where("id = ?", phone.PhoneID).First(&phone).Error; err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Phone number already exists",
		})
	}

	/* Query */
	if err := r.DB.Create(&phone).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Favorite Number has been add",
	})
}

func (r *Repository) DeleteFavorites(c *fiber.Ctx) error {
	/* Get Param ID */
	id, _ := strconv.Atoi(c.Params("id"))
	phone := models.FavoriteNumber{
		Id: uint(id),
	}

	/* Query Delete */
	delete := r.DB.Delete(&phone)
	if delete.RowsAffected == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Phone number not found !",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Phone number deleted successfully",
	})
}

func (r *Repository) DetailFavorites(c *fiber.Ctx) error {
	/* Getting Params ID */
	id, _ := strconv.Atoi(c.Params("id"))
	phone := models.FavoriteNumber{
		Id: uint(id),
	}

	/* Query Detail Number */
	if err := r.DB.Preload("PhoneNumber").First(&phone).Error; err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Phone number not found !",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": phone,
	})

}
