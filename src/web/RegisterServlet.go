// RegisterServlet
package web

import (
	"logs"
)

/*
业务Servlet 向容器注册
*/
var ServletMap map[string]func(request *HttpRequest, response *HttpResponse) IWebView = make(map[string]func(request *HttpRequest, response *HttpResponse) IWebView) //ServiceInterface)

type RegisterServlet struct {
	//请求的path路径

	//业务接口

}

func (register *RegisterServlet) Register(path string, service func(request *HttpRequest, response *HttpResponse) IWebView) {
	ServletMap[path] = service
	logs.Debug("注册成功", ServletMap[path])
}
func (register *RegisterServlet) Remove(path string) func(request *HttpRequest, response *HttpResponse) IWebView {
	if obj, ok := ServletMap[path]; ok {
		delete(ServletMap, path)
		return obj
	}
	return nil
}
func (register RegisterServlet) GetService(uri string) func(request *HttpRequest, response *HttpResponse) IWebView {
	if obj, ok := ServletMap[uri]; ok {
		return obj
	}
	return nil
}
