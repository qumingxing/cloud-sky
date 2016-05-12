package action

import (
	"common"
	"entity"
	"service"
	"web"
)

var loginService service.LoginService

type LoginAction struct {
	BaseAction
}

func (loginAction *LoginAction) Login(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	userInfo := CreateUserInfo()
	params := new(web.RequestForm)
	if withParamsUserInfo, ok := params.GetRequestParameters(request, userInfo).(*entity.UserInfo); ok {
		success, resultUserInfo := loginService.Login(withParamsUserInfo)
		if success {
			session := request.GetSession()
			session.SetAttribute("userInfo", resultUserInfo)
			return &web.WebView{"success.html", resultUserInfo, "success template"}
		} else {
			return &web.WebView{"success.html", resultUserInfo, "success template"}
		}
	}
	return &web.WebView{"error404.html", nil, "error template"}
}
func (loginAction *LoginAction) LoginPage(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return &web.WebView{"login.html", nil, "login template"}
}
func CreateUserInfo() *entity.UserInfo {
	return &entity.UserInfo{}
}
func (loginAction *LoginAction) GetUserName(request *web.HttpRequest, response *web.HttpResponse) {
	session := request.GetSession()
	userInfo := session.GetAttribute("userInfo")
	if v, ok := userInfo.(*entity.UserInfo); ok {
		json, _ := common.ObjToJson(loginAction.GetResultMap(v.RealName))
		response.Write(json)
	}
}
