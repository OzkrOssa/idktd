package config

import (
	"os"

	vault "github.com/hashicorp/vault/api"
)

type Vault struct {
	Client *vault.Client
}

// Create a new Vault client
func NewVault() (*Vault, error) {

	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	return &Vault{
		Client: client,
	}, nil
}
