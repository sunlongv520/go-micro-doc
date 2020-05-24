package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
  "github.com/micro/go-micro/registry/consul"
	"tools.jtthink.com/ServicesImpl"
	"tools.jtthink.com/Services"
)

func main()  {

	consulReg:=consul.NewRegistry(registry.Addrs("192.168.29.135:8500"))
	mySerivce:=micro.NewService(
		micro.Name("test.jtthink.com"),
		micro.Address(":8001"),
		micro.Registry(consulReg),
		)
	Services.RegisterTestServiceHandler(mySerivce.Server(),new(ServicesImpl.TestService))
	mySerivce.Run()
}
