package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
)

func callAPI(addr string,path string,method string) (string,error)  {
	req,_:=http.NewRequest(method,"http://"+addr+path,nil)
	client:=http.DefaultClient
	res,err:=client.Do(req)
	if err!=nil{
		return "",err
	}
	defer res.Body.Close()
	buf,_:= ioutil.ReadAll(res.Body)
	 return string(buf),nil
}

func main()  {
	consulReg:=consul.NewRegistry(
		registry.Addrs("192.168.29.135:8500"),
	)
	getService,err:=consulReg.GetService("prodservice")
	if err!=nil{
		log.Fatal(err)
	}
	next:=selector.RoundRobin(getService)
	node,_:=next()

	callRes,err:=callAPI(node.Address,"/v1/prods","GET")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(callRes)








}