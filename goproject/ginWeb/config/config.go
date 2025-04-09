package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DB  DBConfig  `toml:"DB"`
	LOG LOGConfig `toml:"LOG"`
}

type DBConfig struct {
	UserName  string `toml:"user_name"`
	Password  string `toml:"password"`
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	DBName    string `toml:"db_name"`
	Protocol  string `toml:"protocol"`
	BatchSize int    `toml:"batch_size"`
}

type LOGConfig struct {
	LOGFile string `toml:"log_file"`
}

var Cfg *Config = &Config{}

func LoadConfig() error {
	curPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current path has error: %s", err.Error())
	}
	if _, err = toml.DecodeFile(curPath+"/config/config.toml", Cfg); err != nil {
		return fmt.Errorf("decoding TOML has error: %s", err.Error())
	}
	return nil
}
