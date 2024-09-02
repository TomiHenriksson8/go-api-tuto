package routes

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"

    "github.com/TomiHenriksson8/go-api-tuto/handlers"
    "github.com/TomiHenriksson8/go-api-tuto/middleware"
)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
    app.Post("/api/register", handlers.Register(db))
    app.Post("/api/login", handlers.Login(db))

    todoGroup := app.Group("/api/todos", middleware.JWTProtected())
    todoGroup.Get("/", handlers.GetTodos(db))
    todoGroup.Post("/", handlers.CreateTodo(db))
    todoGroup.Put("/:id", handlers.SetTodoCompletionStatus(db))
    todoGroup.Delete("/:id", handlers.DeleteTodo(db))
}
