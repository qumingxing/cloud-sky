package action
/*
首页
*/
import (
	"web"
	"common"
	//"service"
	"gopkg.in/mgo.v2/bson"
)

type HomePageAction struct {
	BaseAction
}
//var productService service.ProductService = new (service.ProductService)
func (homeAction *HomePageAction)LoadHomePageData(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	pageSize := common.StringToInt(request.FormValue("pageSize"))
	pageInfo := new(common.PageInfo)
	pageInfo.PageIndex = 1
	pageInfo.PageSize = pageSize
	productPageData := productService.LoadProductByHomePage(pageInfo, bson.M{}, "-id")
	advertPageData := advertService.LoadAdverts(pageInfo, bson.M{})
	categoriesData := categoryService.LoadCategories(pageInfo, bson.M{})
	userToken := request.FormValue("token")
	existsCart := shoppingCartSerevice.FindCartByCondition(bson.M{"userToken":userToken})
	var resultData map[string]interface{} = make(map[string]interface{})
	resultData["products"] = productPageData.Data
	resultData["adverts"] = advertPageData.Data
	resultData["categories"] = categoriesData.Data
	resultData["shoppingcartItemNumb"] = len(existsCart)
	json, _ := common.ObjToJson(resultData)
	response.Write(json)
	return nil
}