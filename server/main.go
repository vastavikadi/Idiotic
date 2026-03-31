package main

import (
	// "encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"idiotic/config"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

type JSONTextResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user,omitempty"`
}

func main() {
	fmt.Print("Idiotic Here!")

	fmt.Println("Connecting to the database...")
	config.Init()


	app := fiber.New(fiber.Config{
		AppName: "Idiotic",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(JSONTextResponse{Message: "Hey, Enjoy your time with Idiotic!"})
	})

	app.Post("/signup", func(c fiber.Ctx) error {
		var user User
		if err := c.Bind().Body(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		fmt.Println("User requested:", user)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Signup successful",
			"user":    fiber.Map{"username": user.Username, "password": user.Password, "email": user.Email},
		})
	})

	app.Post("/login", func(c fiber.Ctx) error {
		var user User
		if err := c.Bind().Body(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		fmt.Println("User requested:", user)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login successful",
			"user":    fiber.Map{"username": user.Username, "password": user.Password, "email": user.Email},
		})
	})

	app.Listen(":8080")
}
