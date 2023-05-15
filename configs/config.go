package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	APP           aPpConfig
	DB            dBConfig
}

type aPpConfig struct {
	Name string
	GenerateDB bool
	URL string
	Cron string
}

type dBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("database.host", "locahost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.APP = aPpConfig{
		Name: viper.GetString("app.name"),
		GenerateDB: viper.GetBool("app.generateDB"),
		URL: viper.GetString("app.url"),
		Cron: viper.GetString("app.cron"),
	}

	cfg.DB = dBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}


	return nil
}

func GetDB() dBConfig {
	return cfg.DB
}


func GetApp() aPpConfig {
	return cfg.APP
}
