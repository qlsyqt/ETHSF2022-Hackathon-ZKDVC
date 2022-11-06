package routers

import "github.com/gin-gonic/gin"
import "issuerserver/api"

func InitRouters(r *gin.Engine) {
	//创建issuer身份，存入后台
	r.POST("/api/v1/identity/create", api.CreateIdentity)
	r.POST("/api/v1/identity/getbyname", api.GetDidByName)
	//holder访问link接口，获取json数据，以展示出二维码
	r.GET("/api/v1/identity/auth", api.Auth)
	//验证钱包端的zk proof & ecdsa proof，并完成创建
	r.POST("/api/v1/identity/callback", api.CallbackAuth)
	//获取issuer的revocation tree中，某个claim的未撤销证明
	r.POST("/api/v1/identity/nonrevocation", api.NonRevocationProof)
	//获取issuer的所有树根
	r.POST("/api/v1/identity/roots", api.RootsOfIdentity)
}
