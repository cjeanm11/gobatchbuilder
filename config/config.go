package config

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        int    `mapstructure:"PORT"`
	DebugMode   bool   `mapstructure:"debug_mode"`
	//...
}
