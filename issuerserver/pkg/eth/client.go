package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"issuerserver/config"
	"sync"
)

var client *ethclient.Client
var once sync.Once

func GetEthClient() *ethclient.Client {
	ensureInit()
	return client
}

func ensureInit() {
	once.Do(func() {
		cfg := config.GetConfig()
		newClient, err := ethclient.Dial(cfg.Blockchain.Rpc)
		if err != nil {
			panic(err)
		}
		client = newClient
	})
}
