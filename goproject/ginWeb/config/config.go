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

func LoadConfig() {
	curPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("error getting current path: %s", err.Error()))
	}
	if _, err = toml.DecodeFile(curPath+"/config/config.toml", Cfg); err != nil {
		panic(fmt.Sprintf("error decoding TOML: %s", err.Error()))
	}
}
