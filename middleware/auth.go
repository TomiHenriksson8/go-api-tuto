package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
			tokenString := c.Get("Authorization")
			if tokenString == "" {
					return c.Status(401).JSON(fiber.Map{"error": "Missing or malformed JWT"})
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
					}
					return jwtSecret, nil
			})

			if err != nil || !token.Valid {
					return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired JWT"})
			}

			claims := token.Claims.(jwt.MapClaims)
			c.Locals("userID", claims["userID"])
			return c.Next()
	}
}
