package api

import (
	"github.com/gin-gonic/gin"
	"issuerserver/models/response"
	"issuerserver/services"
)

func GetDidByName(c *gin.Context) {
	username := c.Request.URL.Query().Get("username")
	did, err := services.GetDidByName(username)
	if err != nil {
		c.JSON(response.COMMON_FAIL, response.ErrorResponse(response.COMMON_FAIL, err))
	}
	c.JSON(response.SUCCESS, response.NewResponseWithPair("did", did))
}
