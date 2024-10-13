package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TomiHenriksson8/go-api-tuto/models"
)

// GetTodos godoc
// @Summary Get list of todos
// @Description Retrieve todos for the logged-in user, sorted by creation date.
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todos [get]
func GetTodos(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(string)
		if !ok || userID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
				"error": "You must be logged in to view your todos. Please register or log in",
			})
		}

		var todos []models.Todo

		objectID, _ := primitive.ObjectIDFromHex(userID)
		filter := bson.M{"userID": objectID}

		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"CreatedAt", -1}})

		collection := db.Collection("todos")
		cursor, err := collection.Find(context.Background(), filter, findOptions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "Error fetching todos",
			})
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo models.Todo
			if err := cursor.Decode(&todo); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
					"error": "Error decoding todos",
				})
			}
			todos = append(todos, todo)
		}

		if len(todos) == 0 {
			return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
				"message": "No todos found",
				"todos":   []models.Todo{},
			})
		}

		return c.Status(fiber.StatusOK).JSON(todos)
	}
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item for the logged-in user.
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "New Todo"
// @Success 201 {object} models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todos [post]
func CreateTodo(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := new(models.Todo)
		userID := c.Locals("userID").(string)

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid request body"})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(map[string]interface{}{"error": "Todo body cannot be empty"})
		}

		objectID, _ := primitive.ObjectIDFromHex(userID)
		todo.UserID = objectID
		todo.CreatedAt = time.Now()

		collection := db.Collection("todos")
		insertResult, err := collection.InsertOne(context.Background(), todo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "Error creating todo",
			})
		}

		todo.ID = insertResult.InsertedID.(primitive.ObjectID)

		return c.Status(201).JSON(todo)
	}
}

// SetTodoCompletionStatus godoc
// @Summary Update the completion status of a todo
// @Description Update the completion status of a specific todo item by its ID.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param completion body models.TodoCompletionRequest true "Completion Status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todos/{id} [patch]
func SetTodoCompletionStatus(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid ID"})
		}

		var body models.TodoCompletionRequest  // Use the named struct here
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid request body"})
		}

		collection := db.Collection("todos")
		filter := bson.M{"_id": objectID}
		update := bson.M{"$set": bson.M{"completed": body.Completed}}

		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "Error updating todo",
			})
		}

		return c.Status(200).JSON(map[string]interface{}{"success": true})
	}
}

// DeleteTodo godoc
// @Summary Delete a todo by ID
// @Description Delete a specific todo item by its ID.
// @Tags todos
// @Param id path string true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /todos/{id} [delete]
func DeleteTodo(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid ID"})
		}

		collection := db.Collection("todos")
		filter := bson.M{"_id": objectID}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "Error deleting todo",
			})
		}

		return c.Status(200).JSON(map[string]interface{}{"success": true})
	}
}
