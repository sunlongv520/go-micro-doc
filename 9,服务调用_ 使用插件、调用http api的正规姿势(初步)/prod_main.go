package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"gomicro.jtthink.com/ProdService"
)

func main(){
	consulReg:=consul.NewRegistry(
		registry.Addrs("192.168.29.135:8500"),

		)
	ginRouter:=gin.Default()
	v1Group:=ginRouter.Group("/v1")
	{
		v1Group.Handle("POST","/prods", func(context *gin.Context) {
			  context.JSON(200,
			  	gin.H{
			  		"data":ProdService.NewProdList(2)})
		})
	}

	server:=web.NewService(
		web.Name("prodservice"),
		//web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
		)
	server.Init()
	server.Run()

}