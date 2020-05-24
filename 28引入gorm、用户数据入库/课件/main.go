package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"tools.jtthink.com/Services"
	"tools.jtthink.com/ServicesImpl"
	_ "tools.jtthink.com/AppInit"
)

func main()  {

	consulReg:=consul.NewRegistry(registry.Addrs("192.168.29.135:8500"))

	mySerivce:=micro.NewService(
		micro.Name("api.jtthink.com.myapp"),
		micro.Address(":8001"),
		micro.Registry(consulReg),
		)
 	Services.RegisterUserServiceHandler(mySerivce.Server(),new(ServicesImpl.UserService))

	mySerivce.Run()
}
