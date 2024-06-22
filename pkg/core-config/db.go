package coreconfig

import coreutils "github.com/amosehiguese/dc/pkg/core-utils"

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
}

func setDatabaseConfig() *databaseConfig {
	var d databaseConfig
	coreutils.MustMapEnv(&d.Host, "DATABASE_HOST")
	coreutils.MustMapEnv(&d.Port, "DATABASE_PORT")
	coreutils.MustMapEnv(&d.User, "DATABASE_USER")
	coreutils.MustMapEnv(&d.Password, "DATABASE_PASSWORD")
	coreutils.MustMapEnv(&d.Name, "DATABASE_NAME")
	coreutils.MustMapEnv(&d.SslMode, "DATABASE_SSLMODE")

	return &d
}
