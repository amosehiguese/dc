package coreconfig

import (
	coreutils "github.com/amosehiguese/dc/pkg/core-utils"
)

type cacheConfig struct {
	Type     string
	Host     string
	Port     string
	Password string
	DbNo     string
}

func setCacheConfig() *cacheConfig {
	var c cacheConfig
	coreutils.MustMapEnv(&c.Type, "CACHE_TYPE")
	coreutils.MustMapEnv(&c.Host, "CACHE_HOST")
	coreutils.MustMapEnv(&c.Port, "CACHE_PORT")
	coreutils.MustMapEnv(&c.Password, "CACHE_PASSWORD")
	coreutils.MustMapEnv(&c.DbNo, "CACHE_DB_NUMBER")

	return &c
}
