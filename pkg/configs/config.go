package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewConfig() fiber.Config {
	return fiber.Config{
		Views:   newTemplateEngine(),
		AppName: "DistroMash v1.0.1",
	}
}

func newTemplateEngine() *html.Engine {
	engine := html.New("./views", ".html")
	return engine
}
