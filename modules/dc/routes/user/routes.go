package user

import (
	user_api "github.com/amosehiguese/dc/modules/dc/api/user"
	"github.com/amosehiguese/dc/modules/dc/middleware"
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func UserRoutes(v1 fiber.Router, config coreconfig.Config, log *zap.Logger) {
	users := v1.Group("/u")

	db := store.GetDBClient()
	rc, _ := store.RedisConn(config, log)

	userHandler := user_api.NewUserHandler(user_api.NewUserRepo(), db, rc, log)

	users.Get("/", middleware.JWTProtected(), userHandler.AllUsers)
	users.Get("/me", middleware.JWTProtected(), userHandler.CurrentUser)
	users.Get("/:id", middleware.JWTProtected(), userHandler.SingleUser)
	users.Post("/", middleware.JWTProtected(), userHandler.CreateUser)
	users.Patch("/modify", middleware.JWTProtected(), userHandler.UpdateUser)
	users.Patch("/modify-password", middleware.JWTProtected(), userHandler.ChangePassword)

}
