package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func init() {
	initialDefaultsEnv()
}

func initialDefaultsEnv() {
	initialDatabaseEnv()
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("local.env") // name of config file (without extension)
	viper.SetConfigType("env")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)        // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	// viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func initialDatabaseEnv() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "comments-microservice")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_PORT", "5432")
}
