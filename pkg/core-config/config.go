package coreconfig

import (
	coreutils "github.com/amosehiguese/dc/pkg/core-utils"
)

type Config struct {
	Env         string
	Server      *serverConfig
	Database    *databaseConfig
	Cache       *cacheConfig
	Security    *securityConfig
	Telemetry   *telemetryConfig
	SMTP        *smtpConfig
	ObjectStore *objectStoreConfig
}

var c Config

func initConfig() *Config {
	c.Server = setServerConfig()
	c.Database = setDatabaseConfig()
	c.Cache = setCacheConfig()
	c.Security = setSecurityConfig()
	c.Telemetry = setTelemetryConfig()
	c.ObjectStore = setObjectStoreConfig()
	c.SMTP = setSmtpConfig()
	coreutils.MustMapEnv(&c.Env, "ENV")

	return &c
}

func GetConfig() *Config {
	if c.Server == nil || c.Database == nil || c.Cache == nil || c.Security == nil || c.Telemetry == nil || c.ObjectStore == nil {
		c = *initConfig()
	}
	return &c
}
