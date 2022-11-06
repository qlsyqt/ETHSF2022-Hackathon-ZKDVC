package api

import (
	"github.com/gin-gonic/gin"
	"issuerserver/models/response"
	"issuerserver/services"
	"strconv"
)

func Auth(c *gin.Context) {
	offerId, err := strconv.ParseInt(c.Request.URL.Query().Get("offerId"), 10, 64)
	if err != nil {
		c.JSON(response.BAD_REQUEST, response.ErrorResponse(response.BAD_REQUEST, err))
		return
	}
	authRequest, err := services.Auth(offerId)
	if err != nil {
		c.JSON(response.BAD_REQUEST, response.ErrorResponse(response.BAD_REQUEST, err))
		return
	}
	c.JSON(200, response.NewResponseWithData(authRequest))
}
