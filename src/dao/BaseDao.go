package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"logs"
	"reflect"
	"strconv"
	"common"
)

var connSource DataSource = GetDataSource("database.config", 5)

type BaseDao struct{}



func (baseDao *BaseDao) NewTransation() Transaction {
	logs.Info("创建一个事务对象")
	return baseDao
}
func (baseDao *BaseDao) Begin() *TransactionConfig {
	conn := connSource.GetConnection()
	tx, _ := conn.Begin()
	logs.Info("打开事务对象")
	return &TransactionConfig{tx, conn}
}
func (baseDao *BaseDao) Commit(config *TransactionConfig) error {
	logs.Info("提交事务对象")
	defer connSource.DesConnection(config.conn)
	return config.Commit()
}
func (baseDao *BaseDao) Rollback(config *TransactionConfig) error {
	logs.Info("回滚事务对象")
	defer connSource.DesConnection(config.conn)
	return config.Rollback()
}

//增加一条记录
func (baseDao *BaseDao) Add(query string, params ...interface{}) (key int, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", query, params)
	stmt, err := conn.Prepare(query)
	if err == nil {
		result, err1 := stmt.Exec(params...)
		err = err1
		if err1 == nil {
			id, _ := result.LastInsertId()
			key = int(id)
		} else {
			logs.Info(err)
		}
	} else {
		logs.Info(err)
	}
	return
}

//增加一条记录(事务)
func (baseDao *BaseDao) AddTx(transactionConfig *TransactionConfig, query string, params ...interface{}) (affectCount int, err error) {
	logs.Info("execute sql:->", query, params)
	stmt, err := transactionConfig.Prepare(query)
	if err == nil {
		result, err1 := stmt.Exec(params...)
		err = err1
		if err1 == nil {
			count, _ := result.RowsAffected()
			affectCount = int(count)
		} else {
			logs.Info(err)
		}
	} else {
		logs.Info(err)
	}
	return
}

//更新记录
func (baseDao *BaseDao) Update(sql string, params ...interface{}) (affectCount int, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	result, err := conn.Exec(sql, params...)
	if err == nil {
		count, _ := result.RowsAffected()
		affectCount = int(count)
	} else {
		logs.Info(err)
	}
	return
}

//更新记录(事务)
func (baseDao *BaseDao) UpdateTx(transactionConfig *TransactionConfig, sql string, params ...interface{}) (affectCount int, err error) {
	logs.Info("execute sql:->", sql, params)
	result, err := transactionConfig.Exec(sql, params...)
	if err == nil {
		count, _ := result.RowsAffected()
		affectCount = int(count)
	} else {
		logs.Info(err)
	}
	return
}

//删除记录
func (baseDao *BaseDao) Delete(sql string, params ...interface{}) (affectCount int, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	result, err := conn.Exec(sql, params...)
	if err == nil {
		count, _ := result.RowsAffected()
		affectCount = int(count)
	} else {
		logs.Info(err)
	}
	return
}

//删除记录(事务)
func (baseDao *BaseDao) DeleteTx(transactionConfig *TransactionConfig, sql string, params ...interface{}) (affectCount int, err error) {
	logs.Info("execute sql:->", sql, params)
	result, err := transactionConfig.Exec(sql, params...)
	if err == nil {
		count, _ := result.RowsAffected()
		affectCount = int(count)
	} else {
		logs.Info(err)
	}
	return
}

//查询返回一个map类型
func (baseDao *BaseDao) SelectOneMap(sql string, params ...interface{}) (result map[string]interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	columns, _ := rows.Columns()
	count := len(columns)
	sliceAdd, sliceValue := newSlice(count)
	temp := 0
	result = make(map[string]interface{}, count)
	for rows.Next() {
		if temp > 1 {
			return nil, errors.New("返回数据大于1条")
		}
		rows.Scan(sliceAdd...)
		result = dataConvertMap(columns, sliceValue)
		temp++
	}

	return result, nil
}

//查询返回一组map类型
func (baseDao *BaseDao) SelectMultMap(sql string, params ...interface{}) (result []map[string]interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	result = make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		oneMap := dataConvertMap(columns, sliceValue)
		result = append(result, oneMap)
	}
	return result, nil
}

//查询返回一组带分页的map
func (baseDao *BaseDao) SelectMultSplitPageMap(pageInfo *common.PageInfo, sql string, params ...interface{}) (result []map[string]interface{}, info *common.PageInfo, err error) {
	rows, err, info := baseDao.basePage(pageInfo, sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	result = make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		oneMap := dataConvertMap(columns, sliceValue)
		result = append(result, oneMap)
	}

	return result, info, err
}

//查询返回一个数组
func (baseDao *BaseDao) SelectOneArray(sql string, params ...interface{}) (result []interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	columns, _ := rows.Columns()
	count := len(columns)
	sliceAdd, sliceValue := newSlice(count)
	temp := 0
	for rows.Next() {
		if temp > 1 {
			return nil, errors.New("返回数据大于1条")
		}
		rows.Scan(sliceAdd...)
		dataConvertArray(count, sliceValue)
		temp++
	}
	return sliceValue, nil
}

//查询返回一个二维数组
func (baseDao *BaseDao) SelectMultArray(sql string, params ...interface{}) (result [][]interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	result = make([][]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		dataConvertArray(count, sliceValue)
		result = append(result, sliceValue)
	}
	return result, nil
}

//查询返回一个带分页的二维数组
func (baseDao *BaseDao) SelectMultSplitPageArray(pageInfo *common.PageInfo, sql string, params ...interface{}) (result [][]interface{}, info *common.PageInfo, err error) {
	rows, err, info := baseDao.basePage(pageInfo, sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	result = make([][]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		dataConvertArray(count, sliceValue)
		result = append(result, sliceValue)
	}
	return result, info, err
}

//查询返回一个结构体
func (baseDao *BaseDao) SelectOneStruct(resultType interface{}, sql string, params ...interface{}) (resultStruct interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}

	columns, _ := rows.Columns()
	count := len(columns)
	sliceAdd, sliceValue := newSlice(count)
	temp := 0
	result := make(map[string]interface{}, count)
	for rows.Next() {
		if temp > 1 {
			return nil, errors.New("返回数据大于1条")
		}
		rows.Scan(sliceAdd...)
		result = dataConvertMap(columns, sliceValue)
		temp++
	}
	if temp == 1 {
		return baseDao.structBind(resultType, result), err
	}
	return nil, err
}

//查询返回一组结构体
func (baseDao *BaseDao) SelectMultStruct(resultType interface{}, sql string, params ...interface{}) (resultStruct []interface{}, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params)
	rows, err := conn.Query(sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	result := make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		oneMap := dataConvertMap(columns, sliceValue)
		result = append(result, oneMap)
	}
	resultArrStruct := make([]interface{}, 0)
	for _, mapObj := range result {
		resultArrStruct = append(resultArrStruct, baseDao.structBind(resultType, mapObj))
	}
	return resultArrStruct, err
}

//查询返回带分页的一组结构体
func (baseDao *BaseDao) SelectSplitPageStruct(pageInfo *common.PageInfo, resultType interface{}, sql string, params ...interface{}) (info *common.PageInfo, resultStruct []interface{}, err error) {
	rows, err, info := baseDao.basePage(pageInfo, sql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	if rows == nil {
		return
	}
	result := make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	count := len(columns)
	for rows.Next() {
		sliceAdd, sliceValue := newSlice(count)
		rows.Scan(sliceAdd...)
		oneMap := dataConvertMap(columns, sliceValue)
		result = append(result, oneMap)
	}
	resultArrStruct := make([]interface{}, 0)
	for _, mapObj := range result {
		resultArrStruct = append(resultArrStruct, baseDao.structBind(resultType, mapObj))
	}
	return info, resultArrStruct, err
}

//结构体绑定
func (baseDao *BaseDao) structBind(resultType interface{}, mapObj map[string]interface{}) interface{} {
	logs.Info("structBind->", resultType, mapObj)
	newResult := reflect.New(reflect.TypeOf(resultType).Elem())
	elem := newResult.Elem()
	typeOf := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		name := typeOf.Field(i).Name
		switch elem.Field(i).Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			func() {
				if mapObj[name] != nil {
					if v, ok := mapObj[name].(int64); ok {
						elem.Field(i).SetInt(v)
					} else if v, ok := mapObj[name].(string); ok {
						str, _ := strconv.Atoi(v)
						elem.Field(i).SetInt(int64(str))
					}
				}
			}()
		case reflect.Float32, reflect.Float64:
			func() {
				if mapObj[name] != nil {
					if v, ok := mapObj[name].(float64); ok {
						elem.Field(i).SetFloat(v)
					} else if v, ok := mapObj[name].(string); ok {
						str, _ := strconv.ParseFloat(v, 0)
						elem.Field(i).SetFloat(str)
					}
				}
			}()
		case reflect.String:
			func() {
				if v, ok := mapObj[name].(string); ok {
					elem.Field(i).SetString(v)
				} else if v, ok := mapObj[name].(int64); ok {
					elem.Field(i).SetString(strconv.Itoa(int(v)))
				}
			}()
		default:
			func() {
				logs.Info("未找到类型->", name, elem.Field(i).Kind(), mapObj[name])
			}()
		}
	}
	return newResult.Interface()
}

//基础查询带分页
func (baseDao *BaseDao) basePage(pageInfo *common.PageInfo, sql string, params ...interface{}) (rows *sql.Rows, err error, info *common.PageInfo) {
	info = new(common.PageInfo)
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("execute sql:->", sql, params, pageInfo)
	newSql := "select count(1) from(" + sql + ") t"
	count, err := baseDao.Count(newSql, params...)
	if err != nil {
		logs.Info(err)
		return
	}
	fmt.Println(count, err)
	if count > 0 {
		var sumPage int
		if count%pageInfo.PageSize == 0 {
			sumPage = count / pageInfo.PageSize
		} else {
			sumPage = count/pageInfo.PageSize + 1
		}
		if pageInfo.PageIndex <= 0 || ((pageInfo.PageIndex-1)*pageInfo.PageSize) > count {
			pageInfo.PageIndex = 1
		}
		info.PageIndex = pageInfo.PageIndex
		info.PageSize = pageInfo.PageSize
		info.SumCount = count
		info.SumPage = sumPage
		fmt.Println(info)
		querySQL := "select * from (" + sql + ") t limit " + fmt.Sprint(((pageInfo.PageIndex - 1) * pageInfo.PageSize)) + "," + fmt.Sprint(pageInfo.PageSize)
		fmt.Println(querySQL, params)
		rows, err = conn.Query(querySQL, params...)
		fmt.Println(rows, err)
	}
	return
}

//返回一个单一的数量值
func (baseDao *BaseDao) Count(sql string, params ...interface{}) (count int, err error) {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	logs.Info("Count method execute sql:->", sql, params)
	row := conn.QueryRow(sql, params...)
	err = row.Scan(&count)
	return count, err
}

//创建二个数组，一个用于存放值，一个用于存放地址
func newSlice(count int) (valuePtrs []interface{}, values []interface{}) {
	values = make([]interface{}, count)
	valuePtrs = make([]interface{}, count)
	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	return
}

//数组数据类型转换
func dataConvertArray(count int, values []interface{}) {
	var v interface{}
	for i := 0; i < count; i++ {
		val := values[i]
		if b, ok := val.([]byte); ok {
			v = string(b)
		} else {
			v = val
		}
		values[i] = v
	}
}

//map数据类型转换
func dataConvertMap(cols []string, values []interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{}, len(cols))
	var v interface{}
	for i := 0; i < len(cols); i++ {
		val := values[i]
		if b, ok := val.([]byte); ok {
			v = string(b)
		} else {
			v = val
		}
		resultMap[cols[i]] = v
	}
	return resultMap
}
