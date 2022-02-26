package main

import (
	"gin_websocket/conf"
	"gin_websocket/router"
	"gin_websocket/service"
)

func main()  {
	conf.Init()
	go service.Manager.Start()
	r:=router.NewRouter()
	_ =r.Run(conf.HttpPort)
}