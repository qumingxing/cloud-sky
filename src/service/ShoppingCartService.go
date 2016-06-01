package service

import (
	"entity"
	"logs"
	"common"
	"gopkg.in/mgo.v2/bson"
)

type ShoppingCartService struct {

}

const cart_collection = "cart"

var productService ProductService

func (shoppingCartService *ShoppingCartService)AddCart(cart *entity.ShoppingCart) bool {
	flag, err := baseDao.Add(cart_collection, cart)
	if err == nil {
		return flag
	} else {
		logs.Error("添加购物车失败->", err.Error())
	}
	return false
}
func (shoppingCartService *ShoppingCartService)UpdateCartStatus(selector bson.M, obj bson.M) bool {
	flag, err := baseDao.UpdateBySelector(cart_collection, selector, obj)
	if err == nil {
		return flag
	} else {
		logs.Error("更新购物车失败->", err.Error())
	}
	return false
}
func (shoppingCartService *ShoppingCartService)FindCartListPage(pageInfo *common.PageInfo, query bson.M) map[string]interface{} {
	dataMap := make([]map[string]interface{}, 10)
	resultMap := make(map[string]interface{}, 10)
	//pageData := baseDao.FindQueryForPage(cart_collection, query, &dataArr, pageInfo)
	baseDao.FindQuery(cart_collection, query, &dataMap)
	var totalAmount float64 = 0.0
	//总单价
	var amount float64 = 0.0
	//总运费
	var freight float64 = 0.0
	for _, value := range dataMap {
		if productId, ok := value["productId"].(string); ok {
			product := productService.GetProduct(productId)
			value["image"] = product.Image
			value["name"] = product.Name
			value["unitPrice"] = product.RetailPrice
			if qty, ok1 := value["qty"].(int); ok1 {
				amount += product.RetailPrice * float64(qty)
			}
			freight += product.Freight
		}
	}
	totalAmount = amount + freight
	if len(dataMap) > pageInfo.PageSize {
		resultMap["items"] = dataMap[pageInfo.PageIndex - 1:pageInfo.PageSize]//逻辑分页
	} else {
		resultMap["items"] = dataMap[pageInfo.PageIndex - 1:]//逻辑分页
	}
	resultMap["amount"] = amount//shoppingCartSerevice.FindGroupBuyAmount(userToken)
	resultMap["totalAmount"] = totalAmount
	resultMap["freight"] = freight
	return resultMap
}
func (shoppingCartService *ShoppingCartService)FindGroupBuyAmount(userToken string) float64 {
	amountMap := baseDao.FindGroup(cart_collection, bson.D{{"userToken", 1}}, bson.D{{"userToken", userToken}},
		"function(cur,result){result.total+=cur.qty;result.count++;}",
		bson.D{{"total", 0}, {"count", 0}})
	if ccc, ok1 := amountMap[0].(map[string]interface{}); ok1 {
		if total, ok2 := ccc["total"].(float64); ok2 {
			return total
		}
	}
	return 0
}
func (shoppingCartService *ShoppingCartService)FindCartListCount(query bson.M) int {
	cartCount := baseDao.FindQueryForCount(cart_collection, query)
	return cartCount
}
func (shoppingCartService *ShoppingCartService)FindCartByCondition(query bson.M) []entity.ShoppingCart {
	data := make([]entity.ShoppingCart, 10)
	baseDao.FindQuery(cart_collection, query, &data)
	return data
}