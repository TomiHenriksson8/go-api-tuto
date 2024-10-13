package handlers

import (
	"context"
	"time"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/TomiHenriksson8/go-api-tuto/models" // Import models package
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Register godoc
// @Summary Register a new user
// @Description Register a new user by providing a username and password. The password is hashed before storing.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func Register(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User) // Use models.User
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Cannot parse JSON"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{"error": "Could not hash password"})
		}
		user.Password = string(hashedPassword)

		user.ID = primitive.NewObjectID()
		collection := db.Collection("users")
		_, err = collection.InsertOne(context.Background(), user)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{"error": "Could not register user"})
		}

		return c.Status(201).JSON(map[string]interface{}{"message": "User registered"})
	}
}

// Login godoc
// @Summary Login a user and return a JWT token
// @Description Authenticate a user by validating the username and password, then return a JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func Login(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data models.LoginRequest // Use models.LoginRequest

		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Cannot parse JSON"})
		}

		collection := db.Collection("users")
		var user models.User
		err := collection.FindOne(context.Background(), bson.M{"username": data.Username}).Decode(&user)
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid username or password"})
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{"error": "Invalid username or password"})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID":   user.ID.Hex(),
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{"error": "Could not generate token"})
		}

		return c.JSON(map[string]interface{}{"token": tokenString})
	}
}
