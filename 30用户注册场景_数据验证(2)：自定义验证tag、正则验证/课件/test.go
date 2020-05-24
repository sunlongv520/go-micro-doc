package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"tools.jtthink.com/AppLib"
)

type Users struct {
	Username string `validate:"required,min=6,max=20" vmsg:"用户名必须6位以上"`
	Userpwd string `validate:"required,min=6,max=18" vmsg:"用户密码必须6位以上"`
	Testname string `validate:"required,username"  vmsg:"用户名规则不正确"`
}
func main()  {

	 user:=&Users{Username:"shenyi",Userpwd:"123123",Testname:"acdefb"}
     valid:=validator.New()
     //加入自定义的正则验证tag
	 _=AppLib.AddRegexTag("username","[a-zA-Z]\\w{5,19}",valid)
     err:=AppLib.ValidErrMsg(user,valid.Struct(user))
     if err!=nil{
		 log.Fatal(err)
	 }
     fmt.Println("验证成功")
}
