package handlers

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "os"
    "github.com/TomiHenriksson8/go-api-tuto/models"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Register a new user
func Register(db *mongo.Database) fiber.Handler {
    return func(c *fiber.Ctx) error {
        user := new(models.User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
        }

        // Hash the password
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
        }
        user.Password = string(hashedPassword)

        // Insert the user into the database
        user.ID = primitive.NewObjectID()
        collection := db.Collection("users")
        _, err = collection.InsertOne(context.Background(), user)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not register user"})
        }

        return c.Status(201).JSON(fiber.Map{"message": "User registered"})
    }
}

// Login a user and return a JWT token
func Login(db *mongo.Database) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var data struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
        }

        collection := db.Collection("users")
        var user models.User
        err := collection.FindOne(context.Background(), bson.M{"username": data.Username}).Decode(&user)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid username or password"})
        }

        // Compare the provided password with the stored hashed password
        err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid username or password"})
        }

        // Generate a JWT token
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "userID":   user.ID.Hex(),
            "username": user.Username,
            "exp":      time.Now().Add(time.Hour * 72).Unix(),
        })

        tokenString, err := token.SignedString(jwtSecret)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
        }

        return c.JSON(fiber.Map{"token": tokenString})
    }
}
