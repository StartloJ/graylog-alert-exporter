package handlers

import (
	"graylog-alert-exporter/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// GetAlerts will return all alerts in database
func GetAlerts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"alerts": database.GetAllAlerts(),
		},
	})
}

// UpdateAlert will update alert data in database
func UpdateAlert(c *fiber.Ctx) error {
	alert := database.Alert{}
	if err := c.BodyParser(&alert); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"data": fiber.Map{
				"message": err,
			},
		})
	}

	if err := database.InsertAlert(alert); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "resource updated successfully",
		},
	})
}

// DeleteAlert will delete alert by id
func DeleteAlert(c *fiber.Ctx) error {
	logrus.Debugf("Alert id to delete from database: %s", c.Params("id"))
	if err := database.RemoveAlert(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"data": fiber.Map{
				"message": err,
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "resource deleted successfully",
		},
	})
}
