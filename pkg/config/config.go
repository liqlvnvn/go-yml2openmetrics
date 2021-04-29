package config

import "github.com/spf13/viper"

type Config struct {
	InputSource string `mapstructure:"yml_input_source"`
	HostName    string `mapstructure:"url"`
	Port        string `mapstructure:"port"`
	MetricsPath string `mapstructure:"metrics_path"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
