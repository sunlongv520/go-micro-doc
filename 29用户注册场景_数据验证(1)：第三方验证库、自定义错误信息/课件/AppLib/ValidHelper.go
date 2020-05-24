package AppLib

import (
	"fmt"
	"reflect"
)

func  GetValidMsg(obj interface{},field string) {
	getObj := reflect.TypeOf(obj)
	if f,exist:=getObj.Elem().FieldByName(field);exist{
		fmt.Println(f.Tag.Get("vmsg"))
	}
}