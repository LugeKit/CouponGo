package main

import (
	"coupon/conf"
	"coupon/data"
	"coupon/router"
	"net/http"
)

func main() {
	conf.Init()
	data.Init()
	server := router.NewServer()
	httpServer := http.Server{
		Addr:    conf.AppConfig.Server.IPAddress,
		Handler: server,
	}
	httpServer.ListenAndServe()
}
