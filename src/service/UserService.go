package service

import (
	"logs"
	"entity"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {

}

const user_collection = "user"

var UserMap map[string]*entity.User = make(map[string]*entity.User, 10)

func (userService *UserService)AddUser(order *entity.User) bool {
	flag, err := baseDao.Add(user_collection, order)
	if err == nil {
		return flag
	} else {
		logs.Error("添加用户失败->", err.Error())
	}
	return false
}
func (userService *UserService)UpdateUser(selector bson.M, obj bson.M) bool {
	flag, err := baseDao.UpdateBySelector(user_collection, selector, obj)
	if err == nil {
		return flag
	} else {
		logs.Error("更新用户失败->", err.Error())
	}
	return false
}
func (userService *UserService)GetUser(query bson.M) *entity.User {
	var user[]*entity.User = make([]*entity.User, 1)
	baseDao.FindQuery(user_collection, query, &user)
	if len(user) > 0 {
		return user[0]
	}
	return nil
}
func (userService *UserService)SaveCacheUser(userToken string, user *entity.User) {
	UserMap[userToken] = user
}
func (userService *UserService)GetCacheUser(userToken string) *entity.User {
	return UserMap[userToken]
}