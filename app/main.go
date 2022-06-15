package main

import (
  "github.com/misaliperver/golish/routes"
  "github.com/misaliperver/golish/controllers"
  
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  app := fiber.New()
  app.Use(cors.New())

  app.Static("/images", "./tmp")

  
  routes.ApiRoutes(app)
  
  app.Get("/", controllers.AppMainController)

  app.Listen(":3000")
}