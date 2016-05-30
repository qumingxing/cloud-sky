package action

import (
	"service"
	"web"
	"common"
	"entity"
	"gopkg.in/mgo.v2/bson"
)

type CategoryAction struct {
	BaseAction
}

var categoryService service.CategoryService

func (categoryAction *CategoryAction)AddCategory(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	category := new(entity.Category)
	requestForm.GetRequestParameters(request, category)
	objectId := bson.NewObjectId()
	category.PId = objectId
	category.Id = objectId.Hex()
	flag := categoryService.AddCategory(category)
	json, _ := common.ObjToJson(categoryAction.GetResultMap(flag))
	response.Write(json)
	return nil
}
