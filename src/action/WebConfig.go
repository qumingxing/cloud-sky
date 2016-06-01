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
	var orderAction OrderAction
	var userAction UserAction
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
	//获取商品详情
	config[apiPath + "/product/opt/get"] = productAction.GetProduct
	//添加到购物车
	config[apiPath + "/shoppingcart/opt/addto"] = shoppingCartAction.AddCart
	//查看购物车列表
	config[apiPath + "/shoppingcart/opt/info"] = shoppingCartAction.FindShoppingCartList
	//保存订单
	config[apiPath + "/shoppingcart/opt/checkout"] = orderAction.SaveOrder
	//获取订单详情
	config[apiPath + "/order/opt/info"] = orderAction.GetOrder
	//取消订单
	config[apiPath + "/order/opt/cancel"] = orderAction.UpdateOrderStatus
	//发现列表
	config[apiPath + "/search/opt/search"] = productAction.FindProductList
	//我的订单列表
	config[apiPath + "/order/opt/list"] = orderAction.FindOrderList
	//登录
	config[apiPath + "/client/opt/signon"] = userAction.Login
	//注册
	config[apiPath + "/client/opt/signup"] = userAction.RegisterUser
	logs.Info("==============初始化servlet结束================")

}
func GetConfig() map[string]func(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	return config
}
