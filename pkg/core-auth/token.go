package coreauth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Access  string
	Refresh string
}

func GenerateNewToken(id string, credentials []string) (*Token, error) {
	accessToken, err := generateAccessToken(id, credentials)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateAccessToken(id string, credentials []string) (string, error) {
	config := coreconfig.GetConfig().Security
	minCount, err := strconv.Atoi(config.JwtSecretKeyExp)
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minCount)).Unix()
	claims["appointment:create"] = false
	claims["appointment:delete"] = false
	claims["appointment:read"] = false
	claims["appointment:update"] = false
	claims["doctor:create"] = false
	claims["doctor:delete"] = false
	claims["doctor:read"] = false
	claims["doctor:update"] = false
	claims["slot:create"] = false
	claims["slot:delete"] = false
	claims["slot:read"] = false
	claims["slot:update"] = false
	claims["user:create"] = false
	claims["user:delete"] = false
	claims["user:read"] = false
	claims["user:update"] = false

	// creds := make(map[string]bool)
	for _, credential := range credentials {
		claims[credential] = true
	}
	// claims["credentials"] = creds

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateRefreshToken() (string, error) {
	config := coreconfig.GetConfig().Security
	hash := sha256.New()
	refresh := config.JwtRefreshKey + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	hoursCount, err := strconv.Atoi(config.JwtRefreshKeyExp)
	if err != nil {
		return "", err
	}

	expTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expTime

	return t, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
