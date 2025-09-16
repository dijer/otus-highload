package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Database    DatabaseConf   `toml:"database"`
	Replicas    []DatabaseConf `toml:"replica"`
	Server      ServerConf     `toml:"server"`
	Auth        AuthConf       `toml:"auth"`
	RedisConfig RedisConf      `toml:"redis"`
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

type RedisConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DBIndex  int    `toml:"db"`
}

func New(filepath string) (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
