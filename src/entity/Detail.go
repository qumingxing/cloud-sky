package entity

import (
	"fmt"
	"strconv"
)

type Detail struct {
	Id        int
	Year      string
	Month     string
	Day       string
	CerNumber string
	Remark    string
	Borrow    float64
	Pay       float64
	Bor_Pay   string
	Balance   float64
	SubId     int
}

func (detail *Detail) String() string {
	fmt.Println(detail.Borrow)
	return "Id:" + strconv.Itoa(detail.Id) + ",Year:" + detail.Year + ",Month:" + detail.Month + ",Day:" + detail.Day + ",CerNumber:" + detail.CerNumber + ",Remark:" + detail.Remark + ",Borrow:" + fmt.Sprintf("%v",detail.Borrow) + ",Pay:" + fmt.Sprint(detail.Pay) + ",Bor_Pay:" + detail.Bor_Pay + ",Balance:" + fmt.Sprint(detail.Balance) + ",SubId:" + strconv.Itoa(detail.SubId)
}
