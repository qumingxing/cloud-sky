package action

import (
	"logs"
	"web"
)

var config map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView = make(map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView)

func init() {
	logs.Info("==============初始化servlet开始================")
	var defaultServlet DefaultServlet
	var loginServlet LoginAction
	var studentServlet StudentServlet
	config["/aaa.html"] = defaultServlet.DefaultMethod
	config["/loginPage.html"] = loginServlet.LoginPage
	config["/login.html"] = loginServlet.Login
	config["/addStudent.html"] = studentServlet.AddStudent
	config["/addStudentPage.html"] = studentServlet.AddStudentPage
	logs.Info("==============初始化servlet结束================")

}
func GetConfig() map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return config
}
