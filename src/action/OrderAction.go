package action

import (
	"web"
	"common"
	"strings"
	"service"
	"entity"
	"gopkg.in/mgo.v2/bson"
)

type OrderAction struct {

}

var orderService service.OrderService
/*
保存订单
*/
func (orderAction *OrderAction)SaveOrder(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	resultMap := make(map[string]interface{}, 10)
	userToken := request.FormValue("token")
	//需要登录
	if userService.GetCacheUser(userToken) == nil {
		resultMap["code"] = -100
		json, _ := common.ObjToJson(resultMap)
		response.Write(json)
		return nil
	}
	products := request.FormValue("products")
	address := new(entity.Address)
	requestForm.GetRequestParameters(request, address)
	productsArr := strings.Split(products, ",")
	order := new(entity.Order)
	productAmount := 0.0
	productFreightAmount := 0.0
	totalAmount := 0.0
	for _, str := range productsArr {
		productIdAndNum := strings.Split(str, "_")
		product := productService.GetProduct(productIdAndNum[0])
		//累计商品金额=累计商品金额*购买数量
		productAmount += (product.RetailPrice * common.StringToFloat(productIdAndNum[1]))
		//累计运费
		productFreightAmount += product.Freight
		order.Products = append(order.Products, *product)
	}
	totalAmount = productAmount + productFreightAmount
	objectId := bson.NewObjectId()
	order.PId = objectId
	order.OrderId = objectId.Hex()
	order.CreateDate = common.GetCurDate()
	order.Status = "0"
	order.UserId = userService.GetCacheUser(userToken).UserId
	order.Amount = productAmount
	order.Freight = productFreightAmount
	order.TotalAmount = totalAmount
	order.Address = address
	orderService.AddOrder(order)
	for _, str := range productsArr {
		//产生订单后从购物车中移除商品
		productIdAndNum := strings.Split(str, "_")
		shoppingCartSerevice.UpdateCartStatus(bson.M{"productId":productIdAndNum[0], "userToken":userToken}, bson.M{"$set":bson.M{"status":"1"}})
	}
	resultMap["orderId"] = order.OrderId
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}
func (orderAction *OrderAction)GetOrder(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	resultMap := make(map[string]interface{}, 10)
	userToken := request.FormValue("token")
	//需要登录
	if userService.GetCacheUser(userToken) == nil {
		resultMap["code"] = -100
		json, _ := common.ObjToJson(resultMap)
		response.Write(json)
		return nil
	}
	orderId := request.FormValue("orderId")
	order := orderService.GetOrderByOrderId(orderId)
	resultMap["data"] = order
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}
func (orderAction *OrderAction)UpdateOrderStatus(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	resultMap := make(map[string]interface{}, 10)
	userToken := request.FormValue("token")
	//需要登录
	if userService.GetCacheUser(userToken) == nil {
		resultMap["code"] = -100
		json, _ := common.ObjToJson(resultMap)
		response.Write(json)
		return nil
	}
	orderId := request.FormValue("orderId")
	orderService.UpdateOrderStatus(bson.M{"orderId":orderId}, bson.M{"$set":bson.M{"status":"3"}})//取消订单
	resultMap["flag"] = true
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}
func (orderAction *OrderAction)FindOrderList(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	resultMap := make(map[string]interface{}, 10)
	userToken := request.FormValue("token")
	//需要登录
	if userService.GetCacheUser(userToken) == nil {
		resultMap["code"] = -100
		json, _ := common.ObjToJson(resultMap)
		response.Write(json)
		return nil
	}
	pageSize := common.StringToInt(request.FormValue("pageSize"))
	pageInfo := new(common.PageInfo)
	pageInfo.PageIndex = 1
	pageInfo.PageSize = pageSize
	orderPageData := orderService.FindOrderList(pageInfo, bson.M{"userId":userService.GetCacheUser(userToken).UserId}, "-id")
	resultMap["data"] = orderPageData.Data
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}