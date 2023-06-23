package main

import (
	"github.com/gofiber/fiber/v2"

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
// @BasePath /api/v1
func main() {
	config := configs.NewConfig()
	app := fiber.New(config)
	// app.Static("/static", "./public")

	routes.PublicRoutes(app)
	routes.SwaggerRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
