package auth_api

import (
	user_api "github.com/amosehiguese/dc/modules/dc/api/user"
	"github.com/amosehiguese/dc/modules/dc/services"
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AuthHandler struct {
	store.DB
	user_api.IUser
	*store.Redis[AuthHandler]
	*zap.Logger
	*coreconfig.Config
}

func NewAuthHandler(u user_api.IUser, db store.DB, rc *redis.Client, log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		DB:     db,
		Redis:  (*store.Redis[AuthHandler])(store.NewRedis[AuthHandler](rc)),
		Logger: log,
		IUser:  u,
		Config: coreconfig.GetConfig(),
	}
}

func (a *AuthHandler) ApiSendVerificationEmail(name, email, verificationToken, origin string) error {
	nve := services.NewVerificationEmail(name, email, verificationToken, origin)
	err := nve.SendVerificationEmail()
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthHandler) ApiSendResetPasswordEmail(name, email, token, origin string) error {
	rpe := services.NewResetPasswordEmail(name, email, token, origin)
	err := rpe.SendResetPasswordEmail()
	if err != nil {
		return err
	}

	return nil
}
