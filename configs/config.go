package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	APP aPpConfig
	DB  dBConfig
}

type aPpConfig struct {
	Name               string
	GenerateDB         bool
	Cron               string
	URL_WORDPRESS      string
	WP_UPLOADS         string
	URL_PUBLIC_UPLOADS string
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
	viper.SetDefault("app.wp_uploads", "/wp-content/uploads")
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
		Name:               viper.GetString("app.name"),
		GenerateDB:         viper.GetBool("app.generateDB"),
		URL_WORDPRESS:      viper.GetString("app.url_wordpress"),
		URL_PUBLIC_UPLOADS: viper.GetString("app.url_public_uploads"),
		WP_UPLOADS:         viper.GetString("app.wp_uploads"),
		Cron:               viper.GetString("app.cron"),
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
