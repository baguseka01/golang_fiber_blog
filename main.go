package main

import (
	"log"
	"os"

	"github.com/baguseka01/golang_fiber_blog/database"
	"github.com/baguseka01/golang_fiber_blog/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memanggil file .env")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)

}
