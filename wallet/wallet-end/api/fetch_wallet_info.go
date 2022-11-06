package api

import (
	"github.com/gin-gonic/gin"
	"wallet-end/models/response"
	"wallet-end/services"
)

func FetchWalletInfo(c *gin.Context) {
	wallet := services.FetchWalletInfo()
	c.JSON(response.SUCCESS, response.NewResponseWithData(wallet))
}
