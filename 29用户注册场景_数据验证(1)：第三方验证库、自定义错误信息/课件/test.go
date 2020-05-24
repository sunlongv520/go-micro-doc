package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"gopkg.in/go-playground/validator.v9"
	"tools.jtthink.com/AppLib"
)

type Users struct {
	Username string `validate:"required,min=6,max=20"`
	Userpwd string `validate:"required,min=6,max=18" vmsg:"用户密码必须6位以上"`
}
func main()  {

	user:=&Users{Username:"shenyi",Userpwd:"123"}
     valid:=validator.New()
     err:=valid.Struct(user)
     if err!=nil{
		 if errs,ok := err.(validator.ValidationErrors);ok{
			 for _, e := range errs {
				 AppLib.GetValidMsg(user,e.Field())
			 }
		 }


		 log.Fatal(err)
	 }
     fmt.Println("验证成功")
}
