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
func (shoppingCartService *ShoppingCartService)FindCartListPage(pageInfo *common.PageInfo, query bson.M) *common.PageData {
	dataArr := make([]map[string]interface{}, 10)
	pageData := baseDao.FindQueryForPage(cart_collection, query, &dataArr, pageInfo)
	data := pageData.Data
	if arrList, ok := data.(*[]map[string]interface{}); ok {
		for _, value := range *arrList {
			if a, b := value["productId"].(string); b {
				product := productService.GetProduct(a)
				value["image"] = product.Image
				value["name"] = product.Name
			}
		}
	}
	return pageData
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