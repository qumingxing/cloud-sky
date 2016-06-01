package action

import (
	"common"
	"gopkg.in/mgo.v2/bson"
	"entity"
	"web"
	"service"
)

type ShoppingCartAction struct {
	BaseAction
}

var shoppingCartSerevice service.ShoppingCartService

func (shoppingCartAction *ShoppingCartAction)AddCart(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	userToken := request.FormValue("token")
	cart := new(entity.ShoppingCart)
	requestForm.GetRequestParameters(request, cart)
	existsCart := shoppingCartSerevice.FindCartByCondition(bson.M{"userToken":userToken, "productId":cart.ProductId,"status":"0"})
	var resultData map[string]interface{} = make(map[string]interface{})
	resultData["success"] = false
	if len(existsCart) > 0 {
		resultData["msg"] = "该商品已经在购物车中存在!"
	} else {
		objectId := bson.NewObjectId()
		cart.PId = objectId
		cart.Id = objectId.Hex()
		cart.UserToken = userToken
		cart.Status = "0"
		flag := shoppingCartSerevice.AddCart(cart)
		if flag {
			shoppingcartItemNumb := shoppingCartSerevice.FindCartListCount(bson.M{"userToken":userToken,"status":"0"})
			resultData["shoppingcartItemNumb"] = shoppingcartItemNumb
			resultData["success"] = true
		}
	}
	json, _ := common.ObjToJson(resultData)
	response.Write(json)
	return nil
}
func (shoppingCartAction *ShoppingCartAction)FindShoppingCartList(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	pageSize := common.StringToInt(request.FormValue("pageSize"))
	userToken := request.FormValue("token")
	pageInfo := new(common.PageInfo)
	pageInfo.PageIndex = 1
	pageInfo.PageSize = pageSize
	//status 0 表示未生产订单的
	shoppingCartDataMap := shoppingCartSerevice.FindCartListPage(pageInfo, bson.M{"userToken":userToken, "status":"0"})
	resultMap := make(map[string]interface{}, 10)
	resultMap["data"] = shoppingCartDataMap
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}