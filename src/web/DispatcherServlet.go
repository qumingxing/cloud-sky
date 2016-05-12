// test1
package web

import (
	"common"
	"fmt"
	"html/template"
	"logs"
	"net/http"
)

type DispatcherServlet struct {
}

var intercepters map[string]WebIntercepter = make(map[string]WebIntercepter)

//添加web拦截器
func (servlet *DispatcherServlet) AddIntercepter(interURI string, intercepter WebIntercepter) {
	if common.IsNotEmpty(interURI) && common.IsNotBlank(intercepter) {
		intercepters[interURI] = intercepter
		logs.Debug("AddIntercepter", interURI, intercepter)
	}
}
func (servlet *DispatcherServlet) getIntercepter(interURI string) WebIntercepter {
	if common.IsNotEmpty(interURI) {
		return intercepters[interURI]
	}
	return nil
}
func (servlet *DispatcherServlet) processInterceptor(request *HttpRequest, response *HttpResponse) bool {
	uri := request.RequestURI()
	nextProcess := true
	for key, value := range intercepters {
		if common.IsURIMatch(uri, key) {
			nextProcess = value.Intercepter(request, response)
			if !nextProcess {
				break
			}
		}
	}
	return nextProcess
}
func (servlet *DispatcherServlet) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer func() {
		//捕获panic异常防止panic后程序终止运行
		if err := recover(); err != nil {
			logs.Error("发生异常", err)
		}
	}()
	uri := request.RequestURI
	httpResponse := &HttpResponse{response}
	httpRequest := &HttpRequest{httpRequest: request, httpResponse: httpResponse}
	logs.Debug(uri)
	if !servlet.processInterceptor(httpRequest, httpResponse) {
		return
	}
	var register RegisterServlet
	service := register.GetService(uri)
	var webView IWebView
	if service == nil {
		webView = &WebView{"error404.html", nil, "error template"}
	} else {
		webView = service(httpRequest, httpResponse)
	}
	if webView != nil {
		templ := template.New(webView.GetTemplateName())
		logs.Debug(fmt.Sprint(common.GetBasePath(), "\\view\\"))
		templ, _ = template.ParseFiles(webView.GetResultPath(fmt.Sprint(common.GetBasePath(), "\\view\\")))
		if err := templ.Execute(response, webView.GetResultObject()); err != nil {
			logs.Error(err)
		}
	}
}
