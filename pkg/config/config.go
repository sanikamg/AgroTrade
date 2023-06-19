package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	JWT              string `mapstructure:"JWT_CODE"`
	AccountSid       string `mapstructure:"ACCOUNT_SID"`
	AuthToken        string `mapstructure:"AUTH_TOKEN"`
	VerifyServiceSid string `mapstructure:"VERIFY_SERVICE_ID"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "JWT_CODE", "ACCOUNT_SID", "AUTH_TOKEN", "VERIFY_SERVICE_ID",
}

var config Config

func LoadConfig() (Config, error) {

	// Set the configuration search path and file
	//viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("/media/user/OS/My programs/Golang_Project_Ecommerce/")

	// Read in the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	// Bind environment variables

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// Validate the Config struct using the validator package
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

// to get the secret code for jwt
func GetJWTConfig() string {

	return config.JWT
}

func GetTwilioconfig() (string, string, string) {
	return config.AccountSid, config.AuthToken, config.VerifyServiceSid

}
