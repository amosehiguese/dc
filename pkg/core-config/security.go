package coreconfig

import (
	"strings"

	coreutils "github.com/amosehiguese/dc/pkg/core-utils"
)

type securityConfig struct {
	JwtSecretKey     string
	JwtSecretKeyExp  string
	JwtRefreshKey    string
	JwtRefreshKeyExp string
	CorsOrigins      []string
}

func setSecurityConfig() *securityConfig {
	var s securityConfig
	coreutils.MustMapEnv(&s.JwtSecretKey, "JWT_SECRET_KEY")
	coreutils.MustMapEnv(&s.JwtSecretKeyExp, "JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	coreutils.MustMapEnv(&s.JwtRefreshKey, "JWT_REFRESH_KEY")
	coreutils.MustMapEnv(&s.JwtRefreshKeyExp, "JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")

	var coreStr string
	coreutils.MustMapEnv(&coreStr, "CORS_ORIGINS")
	if coreStr != "" {
		s.CorsOrigins = strings.Split(coreStr, ",")
	}

	return &s
}
