package config

import (
	"context"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

type (
	Container struct {
		DB    *DB
		Redis *Redis
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		Name       string
		User       string
		Password   string
	}
	Redis struct {
		Address  string
		Port     string
		Password string
	}
)

// New creates a new configuration container with the values from the environment variables or from Vault
func New(ctx context.Context, vault *vault.Client, viperConf *Viper) (*Container, error) {

	if viperConf.Env != "production" {
		fmt.Println("Loading environment variables")
		db := &DB{
			Connection: os.Getenv("DB_CONNECTION"),
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			Name:       os.Getenv("DB_NAME"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
		}

		redis := &Redis{
			Address: os.Getenv("REDIS_ADDRESS"),
			Port:    os.Getenv("REDIS_PORT"),
		}

		return &Container{
			DB:    db,
			Redis: redis,
		}, nil
	}

	vaultPath := os.Getenv("VAULT_PATH")
	dbSecurePath := os.Getenv("DB_SECURE_PATH")
	redisSecurePath := os.Getenv("REDIS_SECURE_PATH")

	secrets, err := vault.KVv2(vaultPath).Get(ctx, dbSecurePath)
	if err != nil {
		return nil, err
	}

	db := &DB{
		Connection: secrets.Data["connection"].(string),
		Host:       secrets.Data["host"].(string),
		Port:       secrets.Data["port"].(string),
		Name:       secrets.Data["name"].(string),
		User:       secrets.Data["user"].(string),
		Password:   secrets.Data["password"].(string),
	}

	secrets, err = vault.KVv2(vaultPath).Get(ctx, redisSecurePath)
	if err != nil {
		return nil, err
	}

	redis := &Redis{
		Address: secrets.Data["address"].(string),
		Port:    secrets.Data["port"].(string),
	}

	return &Container{
		DB:    db,
		Redis: redis,
	}, nil
}

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
