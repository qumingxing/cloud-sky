package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"common"
	//"fmt"
	"logs"
	"errors"
)

const URL string = "127.0.0.1:27017"

var mgoSession *mgo.Session

type MongodbBaeDao struct {

}

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}
func GetCollection(c string) *mgo.Collection {
	return getSession().DB("shopping").C(c)
}

func (baseDao *MongodbBaeDao)Add(c string, docs interface{}) (bool, error) {
	err := GetCollection(c).Insert(docs)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (baseDao *MongodbBaeDao)FindGroup(c string, groupKeys bson.D, cond bson.D,functionStr string, initial bson.D) []interface{} {
	/*groupResult:=make(map[string]interface{},10)
	getSession().DB("shopping").Run(bson.D{{"group",bson.D{{"ns",c},{"key","userToken"},
		{"$reduce","function(cur,result){result.total+=cur.qty;}"},{"initial",bson.D{{"total",0}}}}}},groupResult)
	fmt.Println(groupResult)
	if aaa,ok:=groupResult["retval"].([]map[string]interface{});ok{
		fmt.Println(aaa)
	}
	if bbb,ok:=groupResult["retval"].([]interface{});ok{
		if ccc,ok1:=bbb[0].(map[string]interface{});ok1{
			fmt.Println(ccc["total"],"--")
		}

	}*/
	groupResult := make(map[string]interface{}, 10)
	getSession().DB("shopping").Run(bson.D{{"group", bson.D{{"ns", c}, {"key", groupKeys},
		{"$reduce", functionStr},{"cond",cond}, {"initial", initial}}}}, groupResult)
	/*fmt.Println(groupResult)
	if aaa,ok:=groupResult["retval"].([]map[string]interface{});ok{
		fmt.Println(aaa)
	}
	if bbb,ok:=groupResult["retval"].([]interface{});ok{
		if ccc,ok1:=bbb[0].(map[string]interface{});ok1{
			fmt.Println(ccc["total"],"--")
		}

	}*/
	if result, ok := groupResult["retval"].([]interface{}); ok {
		return result
	}
	return nil

}
//bo,err:=baseDao.UpdateBySelector("mytest",bson.M{"tname":"guoguo1"},bson.M{"$set":bson.M{"age":2}})
func (baseDao *MongodbBaeDao)UpdateBySelector(c string, selector interface{}, obj interface{}) (bool, error) {
	err := GetCollection(c).Update(selector, obj)
	if err != nil {
		return false, err
	}
	return true, nil
}
//bo,err:=baseDao.Update("mytest","5745350a943c8c1f80f2dfac",bson.M{"age":7})
func (baseDao *MongodbBaeDao)Update(c string, id string, obj interface{}) (bool, error) {
	err := GetCollection(c).UpdateId(bson.ObjectIdHex(id), obj)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (baseDao *MongodbBaeDao)Delete(c string, id string) (bool, error) {
	err := GetCollection(c).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return false, err
	}
	return true, nil
}
//baseDao.GET("mytest","5745350a943c8c1f80f2dfac",stu)
//fmt.Println(stu.Age,stu.Id.Hex())
func (baseDao *MongodbBaeDao)GET(c string, id string, result interface{}) {
	if ok := bson.IsObjectIdHex(id); ok {
		objid := bson.ObjectIdHex(id)
		logs.Debug(objid)
		GetCollection(c).FindId(objid).One(result)
	} else {
		errors.New("无效的Mongodb ")
	}
}
//stus:=make([]Student,10)
//baseDao.FindAll("mytest",&stus)
func (baseDao *MongodbBaeDao)FindAll(c string, result interface{}) {
	GetCollection(c).Find(nil).All(result)
}
func (baseDao *MongodbBaeDao)FindAllForSort(c string, result interface{}, orderby ...string) {
	GetCollection(c).Find(nil).Sort(orderby...).All(result)
}
func (baseDao *MongodbBaeDao)FindForCount(c string, ) int {
	count, _ := GetCollection(c).Find(nil).Count()
	return count
}
//stus:=make([]Student,10)
//baseDao.FindQuery("mytest",bson.M{"tname":"guoguo1"},&stus)
func (baseDao *MongodbBaeDao)FindQuery(c string, query interface{}, result interface{}) {
	GetCollection(c).Find(query).All(result)
}
func (baseDao *MongodbBaeDao)FindQueryForSort(c string, query interface{}, result interface{}, orderby ...string) {
	GetCollection(c).Find(query).Sort(orderby...).All(result)
}
//var stus []Student=[]Student{}
//pageInfo:=&common.PageInfo{PageIndex:1,PageSize:2}
//pageData:= baseDao.FindQueryForPage("mytest", nil, &stus, pageInfo, nil...)
/*if v, ok := pageData.Data.(*[]Student); ok {
for _,b:=range *v{
fmt.Println(b.Id)
}
}*/
//
func (baseDao *MongodbBaeDao)FindQueryForPage(c string, query interface{}, result interface{}, pageinfo *common.PageInfo, orderby ...string) *common.PageData {
	count := baseDao.FindQueryForCount(c, query)
	logs.Debug("分页数量->", count)
	if count > 0 {
		pageinfo.SumCount = count
		if count % pageinfo.PageSize == 0 {
			pageinfo.SumPage = count / pageinfo.PageSize
		} else {
			pageinfo.SumPage = count / pageinfo.PageSize + 1
		}
		queryObject := GetCollection(c).Find(query)
		if len(orderby) > 0 {
			queryObject.Sort(orderby...).Skip((pageinfo.PageIndex - 1) * pageinfo.PageSize).Limit(pageinfo.PageSize).All(result)
		} else {
			queryObject.Skip((pageinfo.PageIndex - 1) * pageinfo.PageSize).Limit(pageinfo.PageSize).All(result)
		}
	}
	pageData := &common.PageData{Pageinfo:pageinfo, Data:result}
	logs.Debug("分页数据集->", pageData.Data, pageData.Pageinfo)
	return pageData
}
func (baseDao *MongodbBaeDao)FindQueryForCount(c string, query interface{}) int {
	count, _ := GetCollection(c).Find(query).Count()
	return count
}