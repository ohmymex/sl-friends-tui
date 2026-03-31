package config

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Token        string        `mapstructure:"token"`
	TokenEncoded bool          `mapstructure:"token_encoded"`
	Filter       string        `mapstructure:"filter"`
	Refresh      time.Duration `mapstructure:"refresh"`
	ShowInternal bool          `mapstructure:"show_internal_names"`
	ShowLindens  bool          `mapstructure:"show_lindens"`
	ShowGroups   bool          `mapstructure:"show_groups"`
	Layout       string        `mapstructure:"layout"`
	Notify       NotifyConfig  `mapstructure:"notify"`
}

type NotifyConfig struct {
	Enabled bool     `mapstructure:"enabled"`
	Users   []string `mapstructure:"users"`
}

func Default() *Config {
	return &Config{
		Filter:      "online",
		Refresh:     5 * time.Second,
		ShowLindens: true,
		ShowGroups:  true,
		Layout:      "dashboard",
	}
}

func LoadFromFile(path string) (*Config, error) {
	v := viper.New()
	setDefaults(v)
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	return unmarshal(v)
}

func LoadWithViper(v *viper.Viper) (*Config, error) {
	setDefaults(v)
	return unmarshal(v)
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("filter", "online")
	v.SetDefault("refresh", 5*time.Second)
	v.SetDefault("show_internal_names", false)
	v.SetDefault("show_lindens", true)
	v.SetDefault("show_groups", true)
	v.SetDefault("layout", "dashboard")
	v.SetDefault("notify.enabled", false)
}

func unmarshal(v *viper.Viper) (*Config, error) {
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return cfg, nil
}

func (c *Config) ResolveToken() (string, error) {
	if c.Token == "" {
		return "", fmt.Errorf("token is required")
	}
	if c.TokenEncoded {
		decoded, err := base64.StdEncoding.DecodeString(c.Token)
		if err != nil {
			return "", fmt.Errorf("decode base64 token: %w", err)
		}
		return string(decoded), nil
	}
	return c.Token, nil
}
