package service

import (
	"common"
	"dao"
	"entity"
	"sql"
)

type LoginService struct {
}

func (login *LoginService) Login(userInfo *entity.UserInfo) (logined bool, resUserInfo *entity.UserInfo) {
	var baseDao dao.BaseDao
	result, err := baseDao.SelectOneStruct(&entity.UserInfo{}, sql.T_USER_LOGIN_SQL, userInfo.UserName)
	if err != nil {
		return false, nil
	}
	if user, ok := result.(*entity.UserInfo); ok {
		//716d78e10adc3949ba59abbe56e057f20f883e
		if user != nil && common.Md5(userInfo.Pwd, "sky") == user.Pwd {
			return true, user
		}
	}
	return
}
