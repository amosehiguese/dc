package coreconfig

import coreutils "github.com/amosehiguese/dc/pkg/core-utils"

type objectStoreConfig struct {
	Bucket         string
	Keyfile        string
	PrivateKey     string
	GoogleAccessID string
	SignUrlExp     string
}

func setObjectStoreConfig() *objectStoreConfig {
	var o objectStoreConfig
	coreutils.MustMapEnv(&o.Bucket, "BUCKET")
	coreutils.MustMapEnv(&o.Keyfile, "KEY_FILE_PATH")
	coreutils.MustMapEnv(&o.PrivateKey, "PRIVATE_KEY")
	coreutils.MustMapEnv(&o.GoogleAccessID, "GOOGLE_ACCESS_ID")
	coreutils.MustMapEnv(&o.SignUrlExp, "SIGNED_URL_EXP")
	return &o
}
