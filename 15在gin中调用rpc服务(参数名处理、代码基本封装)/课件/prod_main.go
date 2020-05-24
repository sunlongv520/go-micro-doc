package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"gomicro.jtthink.com/Weblib"

	"gomicro.jtthink.com/Services"
)

func main(){
	consulReg:=consul.NewRegistry(
		registry.Addrs("192.168.29.135:8500"),
		)

	myService:=micro.NewService(micro.Name("prodservice.client"))
	prodService:=Services.NewProdService("prodservice",myService.Client())
	httpServer:=web.NewService(
		web.Name("httpprodservice"),
	    web.Address(":8001"),
	 	web.Handler(Weblib.NewGinRouter(prodService)),
	    web.Registry(consulReg),
	)



	httpServer.Init()
	httpServer.Run()

}