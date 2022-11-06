package services

import (
	"wallet-end/config"
	"wallet-end/models/response"
	"wallet-end/utils"
)

func FetchWalletInfo() *response.WalletResponse {
	cfg := config.GetConfig()
	resp := response.WalletResponse{
		MainWalletAddress: utils.AddressFromPrivateKey(cfg.PolygonWallet.MainPrivateKey),
		AuxWalletAddress:  utils.AddressFromPrivateKey(cfg.PolygonWallet.AuxPrivateKey),
		Network:           cfg.Blockchain.Network,
	}

	return &resp
}
