package action

import (
	"common"
	"entity"
	"service"
	"web"
)

type SubjectAction struct {
	BaseAction
}

var subjectService *service.SubjectService

func (subjectAction *SubjectAction) SubjectListPage(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var info *common.PageInfo = new(common.PageInfo)
	var subject *entity.Subject
	info.PageIndex = common.StringToInt(request.FormValue("PageIndex"))
	info.PageSize = common.StringToInt(request.FormValue("PageSize"))
	info, data, _ := subjectService.GetSubjectListPage(info, subject)
	return &web.WebView{"subjectList", subjectAction.GetResultPageMap(info, data), "SubjectList"}
}
func (subjectAction *SubjectAction) SubjectList(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var subject *entity.Subject
	data, _ := subjectService.GetSubjectList(subject)
	json, _ := common.ObjToJson(subjectAction.GetResultMap(data))
	response.Write(json)
	return nil
}
func (subjectAction *SubjectAction) AddSubject(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var form web.RequestForm
	var subject *entity.Subject
	obj := form.GetRequestParameters(request, subject)
	if v, ok := obj.(*entity.Subject); ok {
		subject = v
	}
	id := subjectService.SaveSubject(subject)
	json, _ := common.ObjToJson(subjectAction.GetResultMap(id))
	response.Write(json)
	return nil
}
func (subjectAction *SubjectAction) UpdateSubject(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var form web.RequestForm
	var subject *entity.Subject
	obj := form.GetRequestParameters(request, subject)
	if v, ok := obj.(*entity.Subject); ok {
		subject = v
	}
	count := subjectService.UpdateSubject(subject)
	json, _ := common.ObjToJson(subjectAction.GetResultMap(count))
	response.Write(json)
	return nil
}
func (subjectAction *SubjectAction) DeleteSubject(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var subject *entity.Subject
	subject.Id = common.StringToInt(request.FormValue("Id"))
	count := subjectService.DeleteSubject(subject)
	json, _ := common.ObjToJson(subjectAction.GetResultMap(count))
	response.Write(json)
	return nil
}
func (subjectAction *SubjectAction) GetSubject(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	var subject *entity.Subject
	subject.Id = common.StringToInt(request.FormValue("Id"))
	obj := subjectService.GetSubject(subject)
	json, _ := common.ObjToJson(subjectAction.GetResultMap(obj))
	response.Write(json)
	return nil
}
