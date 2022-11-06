package main

import (
	"fmt"
	"wallet-end/config"
	"wallet-end/pkgs/orm"
	"wallet-end/routes"
)

func main() {
	orm.InitMySql()
	cfg := config.GetConfig()
	r := routes.InitRouters()
	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port))
}
