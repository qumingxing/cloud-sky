package action

import (
	"web"
	"logs"
)

type SimpleInterceptor struct {
}

func (interceptor SimpleInterceptor) Intercepter(request *web.HttpRequest, response *web.HttpResponse) bool {
	logs.Info("拦截成功")
	return true
}
