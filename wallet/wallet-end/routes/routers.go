package routes

import (
	"github.com/gin-gonic/gin"
	"wallet-end/api"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/api/identity/fetch", api.FetchIdentity)
	r.GET("/api/wallet/fetch", api.FetchWalletInfo)
	r.GET("/api/badge/fetch", api.FetchBadges)
	r.POST("/api/badge/claim", api.ClaimBadge)
	r.POST("/api/authenticate", api.Authenticate)
	r.GET("/api/claims", api.FetchClaims)
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
