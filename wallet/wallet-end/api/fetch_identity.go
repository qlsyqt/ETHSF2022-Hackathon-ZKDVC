package api

import (
	"github.com/gin-gonic/gin"
	"wallet-end/models/response"
	"wallet-end/services"
)

func FetchIdentity(c *gin.Context) {
	did, err := services.FetchIdentity()
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.COMMON_FAIL, err))
		return
	}
	c.JSON(response.SUCCESS, response.NewResponseWithString(did))
}
