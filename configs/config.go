package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	RabbitMQURL       string `mapstructure:"RABBITMQ_URL"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	cfg.DBDriver = viper.GetString("DB_DRIVER")
	cfg.DBHost = viper.GetString("DB_HOST")
	cfg.DBPort = viper.GetString("DB_PORT")
	cfg.DBUser = viper.GetString("DB_USER")
	cfg.DBPassword = viper.GetString("DB_PASSWORD")
	cfg.DBName = viper.GetString("DB_NAME")
	cfg.RabbitMQURL = viper.GetString("RABBITMQ_URL")
	cfg.WebServerPort = viper.GetString("WEB_SERVER_PORT")
	cfg.GRPCServerPort = viper.GetString("GRPC_SERVER_PORT")
	cfg.GraphQLServerPort = viper.GetString("GRAPHQL_SERVER_PORT")

	return &cfg, nil
}
