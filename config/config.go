package config

import "github.com/spf13/viper"

type Config struct {
	AppEnv               string   `mapstructure:"APP_ENV"`
	AppPort              string   `mapstructure:"APP_PORT"`
	AllowOrigin          string   `mapstructure:"APP_ALLOW_ORIGIN"`
	HomePath             string   `mapstructure:"APP_HOME_PATH"`
	AllowImageExtensions []string `mapstructure:"APP_ALLOW_IMAGE_EXTENSIONS"`
	DBUrl                string   `mapstructure:"DB_URL"`
	JWTSecret            string   `mapstructure:"JWT_SECRET"`
	SessionExpire        int      `mapstructure:"SESSION_EXPIRES"`
	GoogleClientId       string   `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleSecret         string   `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirect       string   `mapstructure:"GOOGLE_REDIRECT"`
	GoogleScopes         []string `mapstructure:"GOOGLE_SCOPE"`
	LoginRedirect        string   `mapstructure:"GOOGLE_REGISTER_REDIRECT"`
	S3Bucket             string   `mapstructure:"AWS_S3_BUCKET_NAME"`
}

func (cfg *Config) IsDevelopment() bool {
	return cfg.AppEnv == "development"
}

func Load(config *Config) error {
	_ = viper.BindEnv("APP_ENV")
	_ = viper.BindEnv("APP_PORT")
	_ = viper.BindEnv("APP_ALLOW_ORIGIN")
	_ = viper.BindEnv("APP_HOME_PATH")
	_ = viper.BindEnv("APP_ALLOW_IMAGE_EXTENSIONS")
	_ = viper.BindEnv("DB_URL")
	_ = viper.BindEnv("JWT_SECRET")
	_ = viper.BindEnv("SESSION_EXPIRES")
	_ = viper.BindEnv("GOOGLE_CLIENT_ID")
	_ = viper.BindEnv("GOOGLE_CLIENT_SECRET")
	_ = viper.BindEnv("GOOGLE_REDIRECT")
	_ = viper.BindEnv("GOOGLE_SCOPE")
	_ = viper.BindEnv("GOOGLE_REGISTER_REDIRECT")
	_ = viper.BindEnv("AWS_S3_BUCKET_NAME")

	viper.AutomaticEnv()
	viper.AllowEmptyEnv(false)

	return viper.Unmarshal(config)
}
