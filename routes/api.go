package routes

import (
  apicontrollers "github.com/misaliperver/golish/controllers/api"
  
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

func ApiRoutes(app *fiber.App) {
  api := app.Group("/api", logger.New())
  
	api.Get("/", apicontrollers.ApiMain)
	api.Get("/:name", apicontrollers.DynamicNameHelloApi)
	api.Post("/upload", apicontrollers.UploadFile)

}