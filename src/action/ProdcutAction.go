package action

import ()
import (
	"common"
	"web"
	"service"
	"entity"
	"gopkg.in/mgo.v2/bson"
	//"fmt"
)

type ProductAction struct {
	BaseAction
}

var productService service.ProductService

func (productAction *ProductAction)AddProduct(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	product := new(entity.Product)
	requestForm.GetRequestParameters(request, product)
	objectId := bson.NewObjectId()
	product.PId = objectId
	product.Id = objectId.Hex()
	flag := productService.AddProdcut(product)
	json, _ := common.ObjToJson(productAction.GetResultMap(flag))
	response.Write(json)
	return nil
}
func (productAction *ProductAction)GetProduct(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	userToken := request.FormValue("token")
	existsCart := shoppingCartSerevice.FindCartByCondition(bson.M{"userToken":userToken})
	productId := request.FormValue("productId")
	product := productService.GetProduct(productId)
	var resultData map[string]interface{} = make(map[string]interface{})
	resultData["data"] = product
	resultData["isCollected"] = true
	resultData["shoppingcartItemNumb"] = len(existsCart)
	json, _ := common.ObjToJson(resultData)
	response.Write(json)
	return nil
}