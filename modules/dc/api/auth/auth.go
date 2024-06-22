package auth_api

import (
	user_api "github.com/amosehiguese/dc/modules/dc/api/user"
	"github.com/amosehiguese/dc/modules/dc/services"
	"github.com/amosehiguese/dc/modules/dc/store"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Auth struct {
	store.DB
	user_api.User
	*store.Redis[Auth]
	*zap.Logger
	*coreconfig.Config
}

func NewAuth(u user_api.User, db store.DB, rc *redis.Client, log *zap.Logger) *Auth {
	return &Auth{
		DB:     db,
		Redis:  (*store.Redis[Auth])(store.NewRedis[Auth](rc)),
		Logger: log,
		User:   u,
		Config: coreconfig.GetConfig(),
	}
}

func (a *Auth) ApiSendVerificationEmail(name, email, verificationToken, origin string) error {
	nve := services.NewVerificationEmail(name, email, verificationToken, origin)
	err := nve.SendVerificationEmail()
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) ApiSendResetPasswordEmail(name, email, token, origin string) error {
	rpe := services.NewResetPasswordEmail(name, email, token, origin)
	err := rpe.SendResetPasswordEmail()
	if err != nil {
		return err
	}

	return nil
}
