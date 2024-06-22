package coreauth

import (
	"strings"

	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenMetadata struct {
	UserID      uuid.UUID
	Credentials map[string]bool
	Exp         int64
}

func ExtractTokenMetadata(c fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		exp := int64(claims["exp"].(float64))

		credentials := map[string]bool{
			// Appointment
			"appointment:create": claims["appointment:create"].(bool),
			"appointment:delete": claims["appointment:delete"].(bool),
			"appointment:read":   claims["appointment:read"].(bool),
			"appointment:update": claims["appointment:update"].(bool),

			// Doctor
			"doctor:create": claims["doctor:create"].(bool),
			"doctor:delete": claims["doctor:delete"].(bool),
			"doctor:read":   claims["doctor:read"].(bool),
			"doctor:update": claims["doctor:update"].(bool),

			// Slot
			"slot:create": claims["slot:create"].(bool),
			"slot:delete": claims["slot:delete"].(bool),
			"slot:read":   claims["slot:read"].(bool),
			"slot:update": claims["slot:update"].(bool),

			// User
			"user:create": claims["user:create"].(bool),
			"user:delete": claims["user:delete"].(bool),
			"user:read":   claims["user:read"].(bool),
			"user:update": claims["user:update"].(bool),
		}

		return &TokenMetadata{
			UserID:      userID,
			Credentials: credentials,
			Exp:         exp,
		}, nil
	}

	return nil, err
}

func extractToken(c fiber.Ctx) string {
	bearer := c.Get("Authorization")
	token := strings.Split(bearer, " ")
	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func verifyToken(c fiber.Ctx) (*jwt.Token, error) {
	config := coreconfig.GetConfig().Security
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
