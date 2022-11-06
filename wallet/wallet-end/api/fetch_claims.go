package api

import (
	"github.com/gin-gonic/gin"
	"wallet-end/models/response"
	"wallet-end/services"
)

func FetchClaims(c *gin.Context) {
	claims, err := services.FetchClaims()
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.COMMON_FAIL, err))
		return
	}
	c.JSON(response.SUCCESS, response.NewResponseWithData(claims))
}
