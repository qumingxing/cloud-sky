package service

import (
	"entity"
	"logs"
	"common"
	"gopkg.in/mgo.v2/bson"
)

type CategoryService struct {

}

const category_collection = "category"

func (categoryService *CategoryService)AddCategory(category *entity.Category) bool {
	flag, err := baseDao.Add(category_collection, category)
	if err == nil {
		return flag
	} else {
		logs.Error("添加分类失败->", err.Error())
	}
	return false
}
func (categoryService *CategoryService)LoadCategories(pageInfo *common.PageInfo, query bson.M) *common.PageData {
	dataArr := make([]entity.Category, 10)
	pageData := baseDao.FindQueryForPage(category_collection, query, &dataArr, pageInfo)
	return pageData
}