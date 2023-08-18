package bootstraps

import "github.com/spf13/viper"

type Env struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
	IsReleaseMode      bool   `mapstructure:"IS_RELEASE_MODE"`
}

func NewEnv(dst string) *Env {
	var env Env

	viper.SetConfigFile(dst)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	return &env
}
