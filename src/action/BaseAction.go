package action

import (
	"common"
)

type BaseAction struct {
	//GetResultMap func(info *common.PageInfo, obj interface{}) map[string]interface{}
}

func (baseAction *BaseAction) GetResultPageMap(info *common.PageInfo, obj interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap["PageIndex"] = info
	resultMap["data"] = obj
	return resultMap
}
func (baseAction *BaseAction) GetResultMap(obj interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{})
	resultMap["data"] = obj
	return resultMap
}
