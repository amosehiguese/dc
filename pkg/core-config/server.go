package coreconfig

import coreutils "github.com/amosehiguese/dc/pkg/core-utils"

type serverConfig struct {
	Port        string
	Address     string
	ReadTimeout string
	Origin      string
}

func setServerConfig() *serverConfig {
	var s serverConfig
	coreutils.MustMapEnv(&s.Address, "SERVER_ADDR")
	coreutils.MustMapEnv(&s.Port, "SERVER_PORT")
	coreutils.MustMapEnv(&s.ReadTimeout, "SERVER_READ_TIMEOUT")
	coreutils.MustMapEnv(&s.Origin, "ORIGIN")

	return &s
}
