package controllers

import (
  "github.com/misaliperver/golish/helpers"
  
  "github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func AppMainController(c *fiber.Ctx) error {
	return c.JSON(helpers.ResponseTransform("Golish Hi!", false))
}
