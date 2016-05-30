package action

import (
	"logs"
	"web"
)

var config map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView = make(map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView)
var apiPath string = "/shopping/api/v1"

func init() {
	logs.Info("==============初始化servlet开始================")
	var defaultServlet DefaultServlet
	var loginServlet LoginAction
	var studentServlet StudentServlet
	var homePageAction HomePageAction
	var productAction ProductAction
	var advertAction AdvertAction
	var categoryAction CategoryAction
	var shoppingCartAction ShoppingCartAction
	config["/aaa.html"] = defaultServlet.DefaultMethod
	config["/loginPage.html"] = loginServlet.LoginPage
	config["/login.html"] = loginServlet.Login
	config["/addStudent.html"] = studentServlet.AddStudent
	config["/addStudentPage.html"] = studentServlet.AddStudentPage
	logs.Info("================商城URL注册==========================")
	config[apiPath + "/home/opt/index"] = homePageAction.LoadHomePageData
	config[apiPath + "/addProduct"] = productAction.AddProduct
	config[apiPath + "/addAdvert"] = advertAction.AddAdvert
	config[apiPath + "/addCategory"] = categoryAction.AddCategory
	config[apiPath + "/product/opt/get"] = productAction.GetProduct
	config[apiPath + "/shoppingcart/opt/addto"] = shoppingCartAction.AddCart
	config[apiPath + "/shoppingcart/opt/info"] = shoppingCartAction.FindShoppingCartList


	logs.Info("==============初始化servlet结束================")

}
func GetConfig() map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return config
}
