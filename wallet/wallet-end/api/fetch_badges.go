package api

import (
	"github.com/gin-gonic/gin"
	"wallet-end/models/response"
	"wallet-end/services"
)

func FetchBadges(c *gin.Context) {
	badges, err := services.FetchBadges()
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.COMMON_FAIL, err))
		return
	}
	c.JSON(response.SUCCESS, response.NewResponseWithData(badges))
}
