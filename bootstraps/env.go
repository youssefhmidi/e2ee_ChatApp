package bootstraps

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	SocketTokenSecret  string `mapstructure:"SOCKET_TOKEN_SECRET"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	SocketTokenExpiry  int    `mapstructure:"SOCKET_TOKEN_EXPIRY"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
	IsReleaseMode      bool   `mapstructure:"IS_RELEASE_MODE"`
}

// dst : the location of the env
func NewEnv(dst string) *Env {
	var env Env

	log.Println("searchig of a .env file...")
	viper.SetConfigFile(dst)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	log.Println(".env file has been found and read from")

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}
	log.Println(".env file has been mapped to the Env Struct")

	return &env
}
