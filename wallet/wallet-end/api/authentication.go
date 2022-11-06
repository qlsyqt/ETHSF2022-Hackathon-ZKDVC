package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iden3/iden3comm/protocol"
	"wallet-end/models/response"
	"wallet-end/services"
)

func Authenticate(c *gin.Context) {
	request := protocol.AuthorizationRequestMessage{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.COMMON_FAIL, err))
		return
	}

	err = services.Authentication(&request)
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.COMMON_FAIL, err))
		return
	}

	c.JSON(response.SUCCESS, response.NewResponseWithSuccess())
}
