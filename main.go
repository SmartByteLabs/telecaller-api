package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/princeparmar/telecaller-app/controller"
	"github.com/rightjoin/fig"
	"github.com/rightjoin/fuel"
)

func main() {
	server := fuel.NewServer()
	server.AddService(&controller.UploaderService{})
	server.Version = "1"
	server.Prefix = "api"
	server.Port = fig.Int("service.port")
	server.Run()
}
