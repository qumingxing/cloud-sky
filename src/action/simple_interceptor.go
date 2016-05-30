package action

import (
	"web"
	"logs"
)

type SimpleInterceptor struct {
}
var header map[string]string = map[string]string{"Access-Control-Allow-Origin":"*","Access-Control-Allow-Headers":"Origin, X-Requested-With, Content-Type, Accept"}
func (interceptor SimpleInterceptor) Intercepter(request *web.HttpRequest, response *web.HttpResponse) bool {
	logs.Info("拦截成功")
	response.WriteHeader(header)
	return true
}
