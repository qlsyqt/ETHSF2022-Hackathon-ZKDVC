package response

type WalletResponse struct {
	MainWalletAddress string `json:"mainWalletAddress"`
	AuxWalletAddress  string `json:"auxWalletAddress"`
	Network           string `json:"network"`
}
