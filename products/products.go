package products

import (
	"github.com/gofiber/fiber"
)

func GetAllProducts(c *fiber.Ctx) {
	c.Send("Lista Todos Os Produtos.\n")
}

func GetProduct(c *fiber.Ctx) {
	c.Send("Produto.\n")
}

func AddProduct(c *fiber.Ctx) {
	c.Send("Novo Produto.\n")
}

func DeleteProduct(c *fiber.Ctx) {
	c.Send("Produto Deletado.\n")
}
