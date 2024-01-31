package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBSource         string `mapstructure:"DB_SOURCE"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	TokenSymetricKey string `mapstructure:"TOKEN_SYMETRIC_KEY"`
	EmailName        string `mapstructure:"EMAIL_NAME"`
	EmailSender      string `mapstructure:"EMAIL_SENDER"`
	EmailPassword    string `mapstructure:"EMAIL_PASSWORD"`
	AWSRegion        string `mapstructure:"AWS_REGION"`
	AWSS3Bucket      string `mapstructure:"AWS_S3_BUCKET"`
	AWSAccessKey     string `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretKey     string `mapstructure:"AWS_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
