package entity

import (
	"strconv"
)

type Subject struct {
	Id        int
	SubCode   string
	SubName   string
	SubRemark string
}

func (subject *Subject) String() string {
	return "Id=" + strconv.Itoa(subject.Id) + ",SubCode:" + subject.SubCode + ",SubName:" + subject.SubName + ",SubRemark:" + subject.SubRemark
}
