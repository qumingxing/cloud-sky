package service

import (
	"entity"
	"dao"
	"logs"
	"common"
	"gopkg.in/mgo.v2/bson"
)

type AdvertService struct {

}

const advert_collection = "advert"

var baseDao dao.MongodbBaeDao

func (advertService *AdvertService)AddAdvert(advert *entity.Adverts) bool {
	flag, err := baseDao.Add(advert_collection, advert)
	if err == nil {
		return flag
	} else {
		logs.Error("添加广告失败->", err.Error())
	}
	return false
}
func (advertService *AdvertService)LoadAdverts(pageInfo *common.PageInfo, query bson.M) *common.PageData {
	dataArr := make([]entity.Adverts, 10)
	pageData := baseDao.FindQueryForPage(advert_collection, query, &dataArr, pageInfo, nil...)
	return pageData
}