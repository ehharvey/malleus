package config

import "github.com/spf13/viper"

func InitializeConfig() {
	// Here lies all of the default values for configs

	// Database: internal/infrastructure/db
	viper.SetDefault("db.driver", "postgresql")
	viper.SetDefault("db.username", "postgresql")
	viper.SetDefault("db.password", "postgresql")
	viper.SetDefault("db.host", "postgresql.default.svc.cluster.local")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.database_name", "malleus")
	viper.SetDefault("db.pool.max_conns", 10)

	// Logging Level
	viper.SetDefault("logging.level", "INFO")

	// Listen port
	viper.SetDefault("server.port", "8080")

	// Here lies where we source config from
	viper.SetEnvPrefix("MALLEUS")

}

func LoadConfig() {
	viper.AutomaticEnv()
}
