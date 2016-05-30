package action

import (
	"service"
	"web"
	"common"
	"entity"
	"gopkg.in/mgo.v2/bson"
)

type AdvertAction struct {
	BaseAction
}

var advertService service.AdvertService
var requestForm web.RequestForm

func (advertAction *AdvertAction)AddAdvert(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	advert := new(entity.Adverts)
	requestForm.GetRequestParameters(request, advert)
	objectId := bson.NewObjectId()
	advert.PId = objectId
	advert.Id = objectId.Hex()
	flag := advertService.AddAdvert(advert)
	json, _ := common.ObjToJson(advertAction.GetResultMap(flag))
	response.Write(json)
	return nil
}
