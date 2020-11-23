package main

import (
	"log"
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
)

func initServer() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// This here will appear as a label, one can also use
	// fiberprometheus.NewWith(servicename, namespace, subsystem )
	prometheus := fiberprometheus.New("hello-world-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Static("/static", "./public")
	app.Use(logger.New())
	app.Use(limiter.New())
	app.Use(cors.New())
	app.Use(requestid.New())

	app.Get("/text", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	type testResponse struct {
		Field1 string `json:"field_1"`
		Field2 string `json:"field_2"`
		Field3 string `json:"field_3"`
	}

	app.Get("/api", func(c *fiber.Ctx) error {
		tr := testResponse{"11111111111111", "22222222222222222", "3333333333333333"}
		return c.JSON(tr)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello World",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	return app
}

func main() {

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "3000"
	}
	app := initServer()
	log.Fatal(app.Listen(":" + port))
}
