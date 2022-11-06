package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iden3/iden3comm/protocol"
	"issuerserver/models/response"
	"issuerserver/services"
)

// Verify user response, and issuer claim, then send claim to client. Client would store it
func CallbackAuth(c *gin.Context) {
	callbackResp := protocol.AuthorizationResponseMessage{}
	err := c.BindJSON(&callbackResp)
	if err != nil {
		c.JSON(response.BAD_REQUEST, response.ErrorResponse(response.BAD_REQUEST, err))
		return
	}

	claimBody, issuer, err := services.CallbackAuth(&callbackResp)
	if err != nil {
		c.JSON(response.SUCCESS, response.ErrorResponse(response.BAD_REQUEST, err))
		return
	}

	payload := make(map[string]string)
	payload["claim"] = claimBody
	payload["issuer"] = issuer
	c.JSON(200, response.NewResponseWithData(payload))
}
