package action
/*
首页
*/
import (
	"web"
	"common"
	"logs"
)

type HomePageAction struct {
	BaseAction
}

func (homeAction *HomePageAction)LoadHomePageData(request *web.HttpRequest, response *web.HttpResponse) web.IWebView{
	var resultData map[string]interface{} = make(map[string]interface{})
	resultData["a"] = "123456"
	resultData["b"]="789123"
	logs.Debug("参数"+request.FormValue("pageSize"))
	json, _ := common.ObjToJson(resultData)
	response.Write(json)
	return nil
}