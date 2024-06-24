package user_api

import (
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserHandler struct {
	store.DB
	IUser
	*store.Redis[UserHandler]
	*zap.Logger
	*coreconfig.Config
}

func NewUserHandler(u IUser, db store.DB, rc *redis.Client, log *zap.Logger) *UserHandler {
	return &UserHandler{
		DB:     db,
		Redis:  (*store.Redis[UserHandler])(store.NewRedis[UserHandler](rc)),
		Logger: log,
		IUser:  u,
		Config: coreconfig.GetConfig(),
	}
}
