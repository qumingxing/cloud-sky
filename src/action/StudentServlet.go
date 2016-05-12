package action

import (
	"web"
)

type StudentServlet struct {
}

func (stu *StudentServlet) LoadAllStudent(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return &web.WebView{"studentList.html", nil, "aaa template"}
}
func (stu *StudentServlet) AddStudentPage(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return &web.WebView{"add.html", nil, "add student template"}
}
func (stuObj *StudentServlet) AddStudent(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	request.Redirect("http://www.baidu.com")
	/*headers := make(map[string]string)
	headers["a"] = "123"
	headers["b"] = "456"
	response.WriteHeader(headers)
	response.Write("---------------------1-----------------")*/
	return nil

}
