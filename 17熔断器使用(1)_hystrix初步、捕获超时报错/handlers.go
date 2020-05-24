package Weblib

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"gomicro.jtthink.com/Services"
)
//gin 的方法部分
func GetProdsList(ginCtx *gin.Context)  {
	 	prodService:=ginCtx.Keys["prodservice"].(Services.ProdService)
		var prodReq Services.ProdsRequest
		err:=ginCtx.Bind(&prodReq)
		if err!=nil{
			ginCtx.JSON(500,gin.H{"status":err.Error()})
		}else{
			//熔断代码改造
			//第一步，配置config
			configA:=hystrix.CommandConfig{
				Timeout:5000,
			}
			//第二步：配置command
			hystrix.ConfigureCommand("getprods",configA)
			//第三步：执行，使用Do方法
			var prodRes *Services.ProdListResponse
			err:=hystrix.Do("getprods", func() error {
				prodRes,err= prodService.GetProdsList(context.Background(),&prodReq)
				return err
			},nil)

			if err!=nil{
				ginCtx.JSON(500,gin.H{"status":err.Error()})
			}else{
				ginCtx.JSON(200,gin.H{"data":prodRes.Data})
			}


		}

}