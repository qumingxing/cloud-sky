package main

import (
	"common"
	"dao"
	"entity"
	"fmt"
	"logs"
	"service"
	"time"
	//"unsafe"
	//"task"
)

const oneDay time.Duration = 20

func main() {
	//http://jishi.woniu.com/9yin/getGameServer.html?serverId=7100045&_=1432199104580

	/*aa := common.IsURIMatch("/a/e.html","/*.html")
	fmt.Println(aa)*/
	bb := common.IsURIMatch("/a/ddddd/cf/e.html","/a/*/c/*")
	fmt.Println(bb)
}

type Area struct {
	Id         int
	AreaName   string
	AreaNumber string
	Unkonw     float64
}

func DeleteDetail_Test() {
	var detail *entity.Detail = new(entity.Detail)
	detail.Id = 1
	var service *service.DetailService
	count := service.DeleteDetail(detail)
	logs.Debug(count)
}
func GetDetail_Test() {
	var detail *entity.Detail = new(entity.Detail)
	detail.Id = 1
	var service *service.DetailService
	obj := service.GetDetail(detail)
	logs.Debug(obj)
}
func UpdateDetail_Test() {
	var detail *entity.Detail = new(entity.Detail)
	detail.Id = 1
	detail.Year = "2016"
	detail.Month = "05"
	detail.Day = "12"
	detail.CerNumber = "1"
	detail.Remark = "abs"
	detail.Borrow = 12.1
	detail.Pay = 0.0
	detail.Bor_Pay = "jie"
	detail.Balance = 12.1
	detail.SubId = 2
	var service *service.DetailService
	count := service.UpdateDetail(detail)
	logs.Debug(count)
}
func SaveDetail_Test() {
	var detail *entity.Detail = new(entity.Detail)
	detail.Year = "2015"
	detail.Month = "05"
	detail.Day = "12"
	detail.CerNumber = "1"
	detail.Remark = "abs"
	detail.Borrow = 12.1
	detail.Pay = 0.0
	detail.Bor_Pay = "jie"
	detail.Balance = 12.1
	detail.SubId = 2
	var service *service.DetailService
	id := service.SaveDetail(detail)
	logs.Debug(id)
}
func GetDetailListPage_Test() {
	var detail *entity.Detail
	var detailService *service.DetailService
	info, dataList, err := detailService.GetDetailListPage(&common.PageInfo{PageIndex: 1, PageSize: 2}, detail)
	if err == nil {
		logs.Debug(info)
		logs.Debug(dataList)
	}
}

//------------------------------------------------------------------------------------------
func GetSubject_Test() {
	var subject *entity.Subject = new(entity.Subject)
	subject.Id = 2
	var service *service.SubjectService
	obj := service.GetSubject(subject)
	logs.Debug(obj)
}
func DeleteSubject_Test() {
	var subject *entity.Subject = new(entity.Subject)
	subject.Id = 1
	var service *service.SubjectService
	count := service.DeleteSubject(subject)
	logs.Debug(count)
}
func UpdateSubject_Test() {
	var subject *entity.Subject = new(entity.Subject)
	subject.SubCode = "1001"
	subject.SubName = "你好"
	subject.SubRemark = "你好-1002"
	subject.Id = 1
	var service *service.SubjectService
	count := service.UpdateSubject(subject)
	logs.Debug(count)
}
func SaveSubject_Test() {
	var subject *entity.Subject = new(entity.Subject)
	subject.SubCode = "1001"
	subject.SubName = "你好"
	subject.SubRemark = "你好-1001"
	var service *service.SubjectService
	id := service.SaveSubject(subject)
	logs.Debug(id)
}
func GetSubjectListPage_Test() {
	var subject *entity.Subject
	var service *service.SubjectService
	info, dataList, err := service.GetSubjectListPage(&common.PageInfo{PageIndex: 1, PageSize: 2}, subject)
	if err == nil {
		logs.Debug(info)
		logs.Debug(dataList)
	}
}

//---------------------------------------------------------------------------------------------------------
//增加
func Add_Test(baseDao dao.BaseDao) {
	count, err := baseDao.Add("insert into area (areaName,areaNumber) values (?,?)", "测试", 1888888888)
	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
}

//增加(事务)
func AddTx_Test(baseDao dao.BaseDao, transaction dao.Transaction, transactionConfig *dao.TransactionConfig) (count int, err error) {
	count, err = transaction.AddTx(transactionConfig, "insert into area (areaName,areaNumber) values (?,?)", "测试", 1888888888)

	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
	return
}

//更新
func Update_Test(baseDao dao.BaseDao) {
	count, err := baseDao.Update("update area set areaNumber=? where id=?", 136666666, 24)
	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
}

//删除
func Delete_Test(baseDao dao.BaseDao) {
	count, err := baseDao.Delete("delete from area where id=?", 21)
	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
}

//查询一个map
func SelectOneMap_Test(baseDao dao.BaseDao) {
	obj, err := baseDao.SelectOneMap("select * from area where id=?", 24)
	if err == nil {
		fmt.Println(obj)
	} else {
		fmt.Println(err)
	}
}

//查询一组map
func SelectMultMap_Test(baseDao dao.BaseDao) {
	obj, err := baseDao.SelectMultMap("select * from area")
	if err == nil {
		fmt.Println(obj)
	} else {
		fmt.Println(err)
	}
}

//查询一组map带分页
func SelectMultSplitPageMap_Test(baseDao dao.BaseDao) {
	obj, info, err := baseDao.SelectMultSplitPageMap(&common.PageInfo{PageIndex: 1, PageSize: 2}, "select * from area")
	if err == nil {
		fmt.Println(obj, info)
	} else {
		fmt.Println(err)
	}
}

//查询一个数组
func SelectOneArray_Test(baseDao dao.BaseDao) {
	obj, err := baseDao.SelectOneArray("select * from area where id=?", 24)
	if err == nil {
		fmt.Println(obj)
	} else {
		fmt.Println(err)
	}
}

//查询一个二维数组
func SelectMultArray_Test(baseDao dao.BaseDao) {
	obj, err := baseDao.SelectMultArray("select * from area")
	if err == nil {
		fmt.Println(obj)
	} else {
		fmt.Println(err)
	}
}

//查询一个二维数组带分页
func SelectMultSplitPageArray_Test(baseDao dao.BaseDao) {
	obj, info, err := baseDao.SelectMultSplitPageArray(&common.PageInfo{PageIndex: 1, PageSize: 2}, "select * from area")
	if err == nil {
		fmt.Println(obj, info)
	} else {
		fmt.Println(err)
	}
}

//返回一个结构体
func SelectOneStruct_Test(baseDao dao.BaseDao) {
	var area *Area = &Area{}
	res, err := baseDao.SelectOneStruct(area, "select id as Id,areaName as AreaName,areaNumber as AreaNumber from area where id=?", 24)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}

//结构体所有列表
func SelectMultStruct_Test(baseDao dao.BaseDao) {
	var area *Area
	ara, err := baseDao.SelectMultStruct(area, "select id as Id,areaName as AreaName,areaNumber as AreaNumber from area")
	if err == nil {
		for _, o := range ara {
			if v, ok := o.(*Area); ok {
				fmt.Println(v.Id, v.AreaName, v.AreaNumber, v.Unkonw)
			}
		}
	} else {
		fmt.Println(ara, err)
	}
}

//结构体所有列表分页
func SelectSplitPageStruct_Test(baseDao dao.BaseDao) {
	var area *Area = &Area{}
	pageIndex := &common.PageInfo{PageIndex: 1, PageSize: 2}
	info, ara, err := baseDao.SelectSplitPageStruct(pageIndex, area, "select id as Id,areaName as AreaName,areaNumber as AreaNumber from area")
	fmt.Println("分页信息", info)
	if err == nil {
		for _, o := range ara {
			if v, ok := o.(*Area); ok {
				fmt.Println(v.Id, v.AreaName, v.AreaNumber, v.Unkonw)
			}
		}
	} else {
		fmt.Println(ara, err)
	}
}
