package util
import (
	"github.com/spf13/viper"
	"log"
)

// Config stores all confuration vars of this application
type Config struct {
	DBHost		string `mapstructure:"DB_HOST"`
	DBPort		string `mapstructure:"DB_PORT"`
	DBUser		string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName		string `mapstructure:"DB_NAME"`
	SSLMode 	string `mapstructure:"SSL_MODE"`
	ServerAddr	string `mapstructure:"SERVER_ADDR"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	log.Printf(config.DBHost)
	return
}