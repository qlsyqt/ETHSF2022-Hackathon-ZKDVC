package config

import (
	"github.com/BurntSushi/toml"
	"sync"
)

type Config struct {
	PolygonWallet PolygonWallet
	Server        Server
	PolygonId     PolygonId
	Blockchain    Blockchain
	Database      map[string]Database
	Zkp           map[string]Zkp
}

type Server struct {
	Address string
	Port    int
}

type PolygonWallet struct {
	MainPrivateKey string
	AuxPrivateKey  string
}

type PolygonId struct {
	PrivateKey          string
	AuthRevocationNonce uint64
}

type Blockchain struct {
	Network       string
	StateContract string
	BadgeContract string
}

type Database struct {
	Server string
	Port   int
	Db     string
	DbUser string
	DbPwd  string
}

type Zkp struct {
	WasmPath string
	ZkeyPath string
	VkeyPath string
}

var config Config
var once sync.Once

func ensureInit() {
	filePath := "config.toml"
	once.Do(func() {
		if _, err := toml.DecodeFile(filePath, &config); err != nil {
			panic(err)
		}
	})

}

func GetConfig() *Config {
	ensureInit()
	return &config
}
