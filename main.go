package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/akakream/DistroMash/pkg/configs"
	"github.com/akakream/DistroMash/pkg/routes"
	"github.com/akakream/DistroMash/pkg/utils"
)

// @title DistroMash API
// @version 1.0
// @description DistroMash meshes your Docker Distribution
// @termsOfService http://swagger.io/terms/
// @contact.name Ahmet Kerem Aksoy
// @contact.email a.aksoy@tu-berlin.de
// @host localhost:3000
func main() {
	err := utils.InitSettings()
	if err != nil {
		log.Fatalf("error while initializing the settings, please check the .env file: %v.\n", err)
	}
	config := configs.NewConfig()
	app := fiber.New(config)
	app.Static("/static", "./static")

	if utils.IsEnvDev() {
		app.Use(cors.New())
	}
	app.Use(logger.New())

	routes.PublicRoutes(app)
	routes.UiRoutes(app)
	routes.SwaggerRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
