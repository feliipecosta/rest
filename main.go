package main

import (
	"github.com/feliipecosta/rest/products"
	"github.com/gofiber/fiber"
)

func start(c *fiber.Ctx) {
	c.Send("GET ON /\n")
}

func setupRoutes(app *fiber.App) {
	// GET
	app.Get("/", start)
	app.Get("/api/v1/product", products.GetAllProducts)
	app.Get("/api/v1/product/:id", products.GetProduct)

	// POST
	app.Post("/api/v1/product", products.AddProduct)

	// DELETE
	app.Delete("/api/v1/product/:id", products.DeleteProduct)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen("40839")
}
