package helpers

import "github.com/gofiber/fiber/v2"

func ResponseTransform (message string, isError bool) fiber.Map {
  status := "success"
  
  if isError {
    status = "error"
  }
  
  return fiber.Map{
    "status": status, "message": message, "data": nil}
  
}