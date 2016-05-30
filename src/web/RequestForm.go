// RequestForm
package web

import (
	"logs"
	"reflect"
	"strconv"
)

type RequestForm struct {
}

func (p *RequestForm) GetRequestParameters(httpRequest *HttpRequest, obj interface{}) interface{} {
	request := httpRequest.GetRequest()
	fieldElem := reflect.ValueOf(obj).Elem()
	t := fieldElem.Type() //type of UserInfo
	for i := 0; i < fieldElem.NumField(); i++ {
		field := fieldElem.Field(i)
		requestVal := request.FormValue(t.Field(i).Tag.Get("bson"))//t.Field(i).Name 修改使用`tag`
		//fmt.Println(t.Field(i).Name, field.Kind(), requestVal)
		switch field.Kind() {
		case reflect.String:
			field.SetString(requestVal)
		case reflect.Int, reflect.Int32, reflect.Int64:
			if requestVal != "" {
				if intVal, err := strconv.Atoi(requestVal); err == nil {
					field.SetInt(int64(intVal))
				}
			}
		case reflect.Float32, reflect.Float64:
			if requestVal != "" {
				if intVal, err := strconv.ParseFloat(requestVal, 0); err == nil {
					field.SetFloat(intVal)
				}
			}
		case reflect.Func:
			logs.Warn("GetRequestParameters->", "Func type")
		default:
			logs.Warn("GetRequestParameters->", "not found attribute type")
		}
	}
	return obj
}
