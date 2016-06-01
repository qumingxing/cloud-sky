package action

import (
	"web"
	"service"
	"entity"
	"gopkg.in/mgo.v2/bson"
	"common"
)

type UserAction struct {
	BaseAction
}

var userService service.UserService

func (userAction *UserAction)RegisterUser(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	userToken := request.FormValue("token")
	user := new(entity.User)
	requestForm.GetRequestParameters(request, user)
	userCheck := userService.GetUser(bson.M{"mobile":user.Mobile, "status":user.Status})
	resultMap := make(map[string]interface{}, 10)
	if userCheck == nil {
		objectId := bson.NewObjectId()
		user.PId = objectId
		user.UserId = objectId.Hex()
		user.Password = common.Md5(user.Password, "")
		user.Status = "0"
		user.UserToken = userToken
		user.CreateDate = common.GetCurDate()
		userService.AddUser(user)
		userService.SaveCacheUser(userToken, user)
		resultMap["code"] = 1
		resultMap["token"] = userToken
		//全局公用信息
		userProfile := make(map[string]interface{}, 10)
		userProfile["userId"] = user.UserId
		resultMap["data"] = userProfile
	} else {
		resultMap["code"] = -1
		resultMap["msg"] = "该帐号已存在!"
	}
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}
func (userAction *UserAction)Login(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	userToken := request.FormValue("token")
	user := new(entity.User)
	requestForm.GetRequestParameters(request, user)
	user.Password = common.Md5(user.Password, "")
	user.Status = "0"
	user = userService.GetUser(bson.M{"mobile":user.Mobile, "password":user.Password, "status":user.Status})
	resultMap := make(map[string]interface{}, 10)
	if user == nil {
		resultMap["code"] = -1
	} else {
		if !common.Equals(user.UserToken, userToken) {
			userService.UpdateUser(bson.M{"userId":user.UserId}, bson.M{"$set":bson.M{"userToken":userToken}})
		}
		resultMap["code"] = 1
		resultMap["token"] = userToken
		//全局公用信息
		userProfile := make(map[string]interface{}, 10)
		userProfile["userId"] = user.UserId
		resultMap["data"] = userProfile
		userService.SaveCacheUser(userToken, user)
	}
	json, _ := common.ObjToJson(resultMap)
	response.Write(json)
	return nil
}

