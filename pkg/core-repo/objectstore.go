package corerepo

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	coreconfig "github.com/amosehiguese/dc/pkg/core-config"
	"github.com/gofiber/fiber/v3/log"
	"go.uber.org/zap"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	SCOPE_READ_ONLY  = "https://www.googleapis.com/auth/devstorage.read_only"
	SCOPE_READ_WRITE = "https://www.googleapis.com/auth/devstorage.read_write"
)

type FileUpload struct {
	FileName    string
	ObjectName  string
	ContentType string
}

type ObjectStore struct {
	client *storage.Client
	bucket string
}

func NewObjectStore(log *zap.Logger, bucketName string) (*ObjectStore, error) {
	ctx := context.Background()

	config := coreconfig.GetConfig()
	objStore := config.ObjectStore

	jsonKey, err := os.ReadFile(objStore.Keyfile)
	if err != nil {
		log.Error("Failed to read key file")
		return nil, err
	}

	creds, err := google.CredentialsFromJSON(ctx, jsonKey, SCOPE_READ_WRITE)
	if err != nil {
		log.Error("Failed to create Google Cloud Credentials from JSON")
		return nil, err
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Error("Failed to create Google Cloud Storage client ")
		return nil, err
	}
	return &ObjectStore{
		client: client,
		bucket: bucketName,
	}, nil
}

func (o *ObjectStore) UploadFile(fu FileUpload, log *zap.Logger) error {
	ctx := context.Background()
	f, err := os.Open(fu.FileName)
	if err != nil {
		log.Sugar().Errorf("Failed to open file %s ", fu.FileName)
		return err
	}
	defer f.Close()
	obj := o.client.Bucket(o.bucket).Object(fu.ObjectName)
	wc := obj.NewWriter(ctx)
	wc.ContentType = fu.ContentType

	if _, err := io.Copy(wc, f); err != nil {
		log.Error("Failed to write data to bucket")
		return err
	}

	if err := wc.Close(); err != nil {
		log.Error("Failed to close writer")
		return err
	}

	return nil
}

func (o *ObjectStore) DownloadFile(fu FileUpload, log *zap.Logger) error {
	ctx := context.Background()
	rc, err := o.client.Bucket(o.bucket).Object(fu.ObjectName).NewReader(ctx)
	if err != nil {
		log.Error("Failed to read object from bucket")
		return err
	}
	defer rc.Close()

	f, err := os.Create(fu.FileName)
	if err != nil {
		log.Error("Failed to create file")
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, rc); err != nil {
		log.Error("Failed to write data to file")
		return err
	}

	return nil
}

func (o *ObjectStore) DeleteObject(objectName string, log *zap.Logger) error {
	ctx := context.Background()
	obj := o.client.Bucket(o.bucket).Object(objectName)
	if err := obj.Delete(ctx); err != nil {
		log.Error("Failed to delete object")
		return err
	}

	return nil
}

func (o *ObjectStore) GetSignedURL(objectName string, duration time.Duration) (string, error) {
	c := coreconfig.GetConfig()
	objStoreConfig := c.ObjectStore

	exp, err := strconv.Atoi(objStoreConfig.SignUrlExp)
	if err != nil {
		log.Error("Failed to parse sign url exp")
		return "", nil
	}

	url, err := storage.SignedURL(
		o.bucket,
		objectName,
		&storage.SignedURLOptions{
			Scheme:         storage.SigningSchemeV4,
			Method:         "GET",
			GoogleAccessID: objStoreConfig.GoogleAccessID,
			PrivateKey:     []byte(objStoreConfig.PrivateKey),
			Expires:        time.Now().Add(time.Duration(exp)),
		},
	)
	if err != nil {
		log.Error("Failed to generate signed url")
		return "", err
	}

	return url, nil
}
