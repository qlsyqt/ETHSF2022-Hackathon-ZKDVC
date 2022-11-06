package config

import (
	"github.com/BurntSushi/toml"
	"sync"
)

type Config struct {
	Database       map[string]Database
	Server         Server
	Schema         Schema
	Dvc            Dvc
	Jobs           map[string]JobSchedule
	Blockchain     Blockchain
	Authentication Authentication
}

type Database struct {
	Server string
	Port   int
	Db     string
	DbUser string
	DbPwd  string
}

type Server struct {
	Address string
	Port    int
}

type Schema struct {
	DefaultHash string
}

type Relayer struct {
	BjjPrivateKey string
	BjjPublicKey  string
}

type Dvc struct {
	Host        string
	ReadTimeout int
}

type Blockchain struct {
	Rpc           string
	StateContract string
	PrivateKey    string
}

type JobSchedule struct {
	Cron string
}

type Authentication struct {
	Callback string
}

const tomlPath = "config.toml"

var cfg Config

func GetConfig() Config {
	var once sync.Once
	once.Do(func() {
		_, err := toml.DecodeFile(tomlPath, &cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
