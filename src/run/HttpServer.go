// MVCServer
package main

import (
	"action"
	"fmt"
	"logs"
	"net/http"
	"web"
)

func main() {
	var basicServlet http.Handler = new(web.DispatcherServlet)
	var register web.RegisterServlet
	var simpleInterceptor action.SimpleInterceptor
	if v, ok := basicServlet.(*web.DispatcherServlet); ok {
		v.AddIntercepter("/", simpleInterceptor) // /*.html
	}
	configMap := action.GetConfig()
	for key, value := range configMap {
		logs.Debug("注册Servlet->", " URI: ", key, " Function: ", value)
		register.Register(key, value)
	}
	http.Handle("/", basicServlet)
	fmt.Println("server started")
	http.ListenAndServe(":8989", nil)
}
