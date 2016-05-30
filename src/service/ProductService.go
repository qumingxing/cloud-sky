package service

import (
	"entity"
	"logs"
	//	"fmt"
	"common"
	"gopkg.in/mgo.v2/bson"
)

type ProductService struct {

}

const product_collection = "product"

func (productService *ProductService)LoadProductByHomePage(pageInfo *common.PageInfo, query bson.M, orderby ...string) *common.PageData {
	dataArr := make([]entity.Product, 10)
	pageData := baseDao.FindQueryForPage(product_collection, query, &dataArr, pageInfo, orderby...)
	return pageData
}
func (productService *ProductService)AddProdcut(product *entity.Product) bool {
	flag, err := baseDao.Add(product_collection, product)
	if err == nil {
		return flag
	} else {
		logs.Error("添加商品失败->", err.Error())
	}
	return false
}
func (productService *ProductService)GetProduct(id string) *entity.Product {
	product := new(entity.Product)
	baseDao.GET(product_collection, id, &product)
	return product
}