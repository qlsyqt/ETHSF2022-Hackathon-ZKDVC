package api

import (
	"github.com/gin-gonic/gin"
	"wallet-end/models/response"
	"wallet-end/services"
)

// Send proof to contract and mint one
func ClaimBadge(c *gin.Context) {
	hIndex := c.Request.URL.Query().Get("hIndex")
	err := services.ClaimBadge(hIndex)
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.BAD_REQUEST, err))
		return
	}
	c.JSON(response.SUCCESS, response.NewResponseWithSuccess())
}
