package configs

import (
	"go.uber.org/config"
	"os"
)

type Config struct {
	Database struct {
		Driver string
		URL    string
	}
	JWT struct {
		SecretKey string
	}
}

var configs Config

func Init() error {
	provider, err := config.NewYAML(config.File(os.Getenv("CONFIG_FILE")))

	if err != nil {
		return err
	}

	if err := provider.Get(config.Root).Populate(&configs); err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return configs
}
