package handlers

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "github.com/TomiHenriksson8/go-api-tuto/models"
)

func GetTodos(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
			var todos []models.Todo
			userID := c.Locals("userID").(string)

			objectID, _ := primitive.ObjectIDFromHex(userID)
			filter := bson.M{"userID": objectID}

			collection := db.Collection("todos")
			cursor, err := collection.Find(context.Background(), filter)
			if err != nil {
					return err
			}
			defer cursor.Close(context.Background())

			for cursor.Next(context.Background()) {
					var todo models.Todo
					if err := cursor.Decode(&todo); err != nil {
							return err
					}
					todos = append(todos, todo)
			}

			return c.JSON(todos)
	}
}


func CreateTodo(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
			todo := new(models.Todo)
			userID := c.Locals("userID").(string)

			if err := c.BodyParser(todo); err != nil {
					return err
			}

			if todo.Body == "" {
					return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
			}

			objectID, _ := primitive.ObjectIDFromHex(userID)
			todo.UserID = objectID

			collection := db.Collection("todos")
			insertResult, err := collection.InsertOne(context.Background(), todo)
			if err != nil {
					return err
			}

			todo.ID = insertResult.InsertedID.(primitive.ObjectID)

			return c.Status(201).JSON(todo)
	}
}

func SetTodoCompletionStatus(db *mongo.Database) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        objectID, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
        }

        type RequestBody struct {
            Completed bool `json:"completed"`
        }

        var body RequestBody
        if err := c.BodyParser(&body); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
        }

        collection := db.Collection("todos")
        filter := bson.M{"_id": objectID}
        update := bson.M{"$set": bson.M{"completed": body.Completed}}

        _, err = collection.UpdateOne(context.Background(), filter, update)
        if err != nil {
            return err
        }

        return c.Status(200).JSON(fiber.Map{"success": true})
    }
}

func DeleteTodo(db *mongo.Database) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        objectID, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
        }

        collection := db.Collection("todos")
        filter := bson.M{"_id": objectID}
        _, err = collection.DeleteOne(context.Background(), filter)
        if err != nil {
            return err
        }

        return c.Status(200).JSON(fiber.Map{"success": true})
    }
}
