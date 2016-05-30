package action

import (
	"common"
)

type PageSize int
type BaseAction struct {
	//GetResultMap func(info *common.PageInfo, obj interface{}) map[string]interface{}
	PageSize
}

func (pageSize PageSize)DefaultPageSize() int {
	return 10
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
