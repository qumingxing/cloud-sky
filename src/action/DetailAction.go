package action

import (
	"common"
	"entity"
	"service"
	"web"
)

type DetailAction struct {
	BaseAction
}

var detailService *service.DetailService

func (detailAction *DetailAction) DetailList(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var info *common.PageInfo = new(common.PageInfo)
	var detail *entity.Detail
	info.PageIndex = common.StringToInt(request.FormValue("PageIndex"))
	info.PageSize = common.StringToInt(request.FormValue("PageSize"))
	info, data, _ := detailService.GetDetailListPage(info, detail)
	return &web.WebView{"detailList", detailAction.GetResultPageMap(info, data), "DetailList"}
}
func (detailAction *DetailAction) AddDetail(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var form web.RequestForm
	var detail *entity.Detail
	obj := form.GetRequestParameters(request, detail)
	if v, ok := obj.(*entity.Detail); ok {
		detail = v
	}
	id := detailService.SaveDetail(detail)
	json, _ := common.ObjToJson(detailAction.GetResultMap(id))
	response.Write(json)
	return nil
}
func (detailAction *DetailAction) UpdateDetail(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var form web.RequestForm
	var detail *entity.Detail
	obj := form.GetRequestParameters(request, detail)
	if v, ok := obj.(*entity.Detail); ok {
		detail = v
	}
	count := detailService.UpdateDetail(detail)
	json, _ := common.ObjToJson(detailAction.GetResultMap(count))
	response.Write(json)
	return nil
}
func (detailAction *DetailAction) DeleteDetail(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var detail *entity.Detail
	detail.Id = common.StringToInt(request.FormValue("Id"))
	count := detailService.DeleteDetail(detail)
	json, _ := common.ObjToJson(detailAction.GetResultMap(count))
	response.Write(json)
	return nil
}
func (detailAction *DetailAction) GetDetail(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var detail *entity.Detail
	detail.Id = common.StringToInt(request.FormValue("Id"))
	obj := detailService.GetDetail(detail)
	json, _ := common.ObjToJson(detailAction.GetResultMap(obj))
	response.Write(json)
	return nil
}
