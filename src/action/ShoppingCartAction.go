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
	existsCart := shoppingCartSerevice.FindCartByCondition(bson.M{"userToken":userToken, "productId":cart.ProductId})
	var resultData map[string]interface{} = make(map[string]interface{})
	resultData["success"] = false
	if len(existsCart) > 0 {
		resultData["msg"] = "该商品已经在购物车中存在!"
	} else {
		objectId := bson.NewObjectId()
		cart.PId = objectId
		cart.Id = objectId.Hex()
		cart.UserToken = userToken
		flag := shoppingCartSerevice.AddCart(cart)
		if flag {
			shoppingcartItemNumb := shoppingCartSerevice.FindCartListCount(bson.M{"userToken":userToken})
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
	shoppingCartPageData := shoppingCartSerevice.FindCartListPage(pageInfo, bson.M{"userToken":userToken})
	dataMap := make(map[string]interface{}, 10)
	dataMap["items"] = shoppingCartPageData.Data
	dataMap["amount"] = shoppingCartSerevice.FindGroupBuyAmount(userToken)
	dataMap["totalAmount"] = 123456
	dataMap["freight"] = 1000

	resultMap := make(map[string]interface{}, 10)
	resultMap["data"] = dataMap
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}