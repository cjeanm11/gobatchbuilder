package config

type Config struct {
	LoggingConfig LoggingConfig `mapstructure:"logging"`
}

type LoggingConfig struct {
	HTTPRequests bool   `mapstructure:"http_requests"`
	Level        string `mapstructure:"level"`
	Format       string `mapstructure:"format"`
	Destination  string `mapstructure:"destination"`
}
