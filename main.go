package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrinhoas7/go_abacatepay/service/abacatepayService"
)

func main() {
	app := fiber.New()

	app.Post("/api/pix", func(c *fiber.Ctx) error {
		var req abacatepayService.ChargeRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("JSON inválido")
		}

		token := os.Getenv("ABACATE_TOKEN") // use dotenv se preferir
		resp, err := abacatepayService.CreatePixCharge(token, req)
		if err != nil {
			return c.Status(500).SendString("Erro ao gerar cobrança")
		}

		return c.Send(resp)
	})

	app.Listen(":3000")
}
