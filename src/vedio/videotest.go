package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"common"
	"strings"
	"sort"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Word     string  `xml:"word"`
	Trans    string `xml:"trans"`
	Phonetic string `xml:"phonetic"`
	Tags     string `xml:"tags"`
	Progress int `xml:"progress"`
}
type Result struct {
	Item [] *Item `xml:"item"`
}

var UnitType []string = []string{}

func (item *Item) toString() string {
	return item.Word + " " + item.Phonetic + " \n" + item.Tags + " \n" + item.Trans + "\n"
}
type Student struct {
	Id bson.ObjectId `bson:"_id"`
	Name string `bson:"tname"`
	Age int `bson:"age"`
}
func main() {

	//parse8000()
	//parsecet()
	//fmt.Println(time.Now().Unix())
	//var baseDao = new (dao.MongodbBaeDao)
	//bo,err:=baseDao.Add("mytest",&Student{Id:bson.NewObjectId(),Name:"guoguo1",Age:1})
	//bo,err:=baseDao.Update("mytest",bson.ObjectIdHex("5745350a943c8c1f80f2dfac"),bson.M{"age":7})
	//bo,err:=baseDao.UpdateBySelector("mytest",bson.M{"tname":"guoguo1"},bson.M{"$set":bson.M{"age":2}})
	//baseDao.GET("mytest","5745350a943c8c1f80f2dfac",stu)
	//baseDao.FindAll("mytest",&stus)
	//baseDao.Delete("mytest","57453251943c8c1cd819e2d5")
	/*baseDao.Delete("mytest","574532d0943c8c1d345bf0af")
	baseDao.Delete("mytest","574534ed943c8c1c88a0b546")
	baseDao.Delete("mytest","5745350a943c8c1f80f2dfac")
	baseDao.Delete("mytest","57453994943c8c1fd09cc920")*/
	//Bdate := "2014-06-24 14:30"//时间字符串
	//fmt.Println(t)
	//fmt.Println(bo,err)
}
func parse8000() {
	for i := 1; i <= 39; i++ {
		UnitType = append(UnitType, "8000-" + strconv.Itoa(i))
	}
	bytes, _ := ioutil.ReadFile("f:/ab.xml")
	var result Result;
	err := xml.Unmarshal(bytes, &result)
	if err == nil {
		for _, str := range UnitType {
			addIndex := 0
			for _, val := range result.Item {

				//if strings.Contains(val.Tags,str) {
				if val.Tags == str {
					addIndex++
					aa, _ := common.Append_CreateFile("F:\\word\\8000\\" + val.Tags + ".txt");
					aa.WriteString(strconv.Itoa(addIndex) + ". " + val.toString())
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}
func (a *Result) Len() int {
	return len(a.Item)
}
func (a *Result) Less(i, j int) bool {
	return a.Item[i].Word < a.Item[j].Word
}
func (a *Result) Swap(i, j int) {
	a.Item[i], a.Item[j] = a.Item[j], a.Item[i]
}
func parsecet() {
	for i := 1; i <= 66; i++ {
		UnitType = append(UnitType, "CET-" + strconv.Itoa(i) + "-")
	}
	bytes, _ := ioutil.ReadFile("f:/aa.xml")
	var result *Result;
	err := xml.Unmarshal(bytes, &result)
	sort.Sort(result)
	if err == nil {
		for _, str := range UnitType {
			addIndex := 0
			for _, val := range result.Item {

				if strings.Contains(val.Tags, str) {
					addIndex++
					aa, _ := common.Append_CreateFile("F:\\word\\CST-test\\" + val.Tags + ".txt");
					aa.WriteString(strconv.Itoa(addIndex) + ". " + val.toString())
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}