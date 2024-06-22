package routes

import (
	auth_route "github.com/amosehiguese/dc/modules/dc/routes/auth"
	"github.com/amosehiguese/dc/modules/dc/routes/user"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func MainRoutesV1(api fiber.Router, config coreconfig.Config, log *zap.Logger) {
	v1 := api.Group("/v1")
	user.UserRoutes(v1, config, log)
	auth_route.AuthRoutes(v1, config, log)
}
