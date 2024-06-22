package auth_route

import (
	auth_api "github.com/amosehiguese/dc/modules/dc/api/auth"
	user_api "github.com/amosehiguese/dc/modules/dc/api/user"
	"github.com/amosehiguese/dc/modules/dc/middleware"
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func AuthRoutes(v1 fiber.Router, config coreconfig.Config, log *zap.Logger) {
	auth := v1.Group("/auth")
	db := store.GetDBClient()

	// todo: handle error
	rc, _ := store.RedisConn(config, log)

	authApi := auth_api.NewAuth(user_api.NewUserRepo(), db, rc, log)

	auth.Post("/signup", authApi.Signup)
	auth.Post("/login", authApi.Login)
	auth.Delete("/logout", middleware.JWTProtected(), authApi.Logout, authenticate)
	auth.Post("/verify-email", authApi.VerifyEmail)
	auth.Post("/reset-password", authApi.ResetPassword)
	auth.Post("/forgot-password", authApi.ForgotPassword)
}
