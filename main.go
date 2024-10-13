package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/swaggo/fiber-swagger"
	_ "github.com/TomiHenriksson8/go-api-tuto/docs"

	"github.com/TomiHenriksson8/go-api-tuto/config"
	"github.com/TomiHenriksson8/go-api-tuto/db"
	"github.com/TomiHenriksson8/go-api-tuto/routes"
)


func main() {
    err := config.LoadEnv()
    if err != nil {
        log.Fatal("Error loading .env file", err)
    }

    db, err := database.Connect()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Client().Disconnect(nil)

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))

		app.Get("/api/swagger/*", fiberSwagger.WrapHandler)

    routes.SetupRoutes(app, db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}
