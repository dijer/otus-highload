package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Database DatabaseConf `toml:"database"`
	Server   ServerConf   `toml:"server"`
	Auth     AuthConf     `toml:"auth"`
}

type DatabaseConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type ServerConf struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type AuthConf struct {
	JWTCookieName  string `toml:"jwt_cookie_name"`
	JWTKey         string `toml:"jwt_key"`
	JWTExpireHours int    `toml:"jwt_expire_hours"`
}

func New(filepath string) (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
