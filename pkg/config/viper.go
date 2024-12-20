package config

import "github.com/spf13/viper"

type Viper struct {
	Env string
}

// Read configuration from a file with viper and return a Viper struct
func NewViper() (*Viper, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Viper{
		Env: viper.GetString("app.env"),
	}, nil

}
