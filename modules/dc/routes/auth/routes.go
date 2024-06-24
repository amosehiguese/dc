package auth_route

import (
	auth_api "github.com/amosehiguese/dc/modules/dc/api/auth"
	user_api "github.com/amosehiguese/dc/modules/dc/api/user"
	"github.com/amosehiguese/dc/modules/dc/middleware"
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func AuthRoutes(v1 fiber.Router, config coreconfig.Config, log *zap.Logger, rc *redis.Client) {
	auth := v1.Group("/auth")
	db := store.GetDBClient()

	authHandler := auth_api.NewAuthHandler(user_api.NewUserRepo(), db, rc, log)

	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)
	auth.Delete("/logout", middleware.JWTProtected(), authHandler.Logout, authenticate)
	auth.Post("/verify-email", authHandler.VerifyEmail)
	auth.Post("/reset-password", authHandler.ResetPassword)
	auth.Post("/forgot-password", authHandler.ForgotPassword)
}
