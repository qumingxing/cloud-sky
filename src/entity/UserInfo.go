package entity

import ()

type UserInfo struct {
	Id int
	UserName string
	RealName string
	Pwd      string
	RoleId   string
	Permission
}
