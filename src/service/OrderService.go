package service

import (
	"entity"
	"logs"
	"gopkg.in/mgo.v2/bson"
	"common"
)

type OrderService struct {

}

const order_collection = "order"

func (orderService *OrderService)AddOrder(order *entity.Order) bool {
	flag, err := baseDao.Add(order_collection, order)
	if err == nil {
		return flag
	} else {
		logs.Error("添加订单失败->", err.Error())
	}
	return false
}
func (orderService *OrderService)GetOrderByOrderId(orderId string) *entity.Order {
	order := new(entity.Order)
	baseDao.GET(order_collection, orderId, &order)
	return order
}
func (orderService *OrderService)UpdateOrderStatus(selector bson.M, obj bson.M) bool {
	flag, err := baseDao.UpdateBySelector(order_collection, selector, obj)
	if err != nil {
		logs.Error("更新订单状态失败", err)
	}
	return flag
}
func (orderService *OrderService)FindOrderList(pageInfo *common.PageInfo, query bson.M, orderby ...string) *common.PageData {
	dataArr := make([]entity.Order, 10)
	pageData := baseDao.FindQueryForPage(order_collection, query, &dataArr, pageInfo, orderby...)
	return pageData
}