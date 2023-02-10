package main

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	godotenv "github.com/joho/godotenv"

	internal "go-oauth/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	app.Get("/auth/login", func(c *fiber.Ctx) error {
		state := c.Query("state")

		if state == "" {
			return c.Status(fiber.StatusBadRequest).SendString("State is not data.")
		}

		url := internal.Oauth2().AuthCodeURL(state)
		fmt.Println("url", c.Redirect(url))
		return c.Redirect(url)
	})
	app.Get("/auth/callback", func(c *fiber.Ctx) error {
		return internal.LoginCallback(c)
	})

	app.Listen(":8001")
}
