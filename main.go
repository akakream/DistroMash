package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akakream/DistroMash/pkg/configs"
	"github.com/akakream/DistroMash/pkg/routes"
	"github.com/akakream/DistroMash/pkg/utils"
)

func main() {
	config := configs.NewConfig()
	app := fiber.New(config)
	// app.Static("/static", "./public")

	routes.PublicRoutes(app)
	routes.SwaggerRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
