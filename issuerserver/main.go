package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"issuerserver/config"
	"issuerserver/pkg/orm"
	"issuerserver/routers"
	"issuerserver/start"
)

func main() {
	//job.StartAutoRevoke()
	orm.Init()
	start.CreateDefaultUser()
	r := gin.Default()
	cfg := config.GetConfig()
	routers.InitRouters(r)
	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port))
}
