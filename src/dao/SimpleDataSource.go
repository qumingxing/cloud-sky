package dao

import (
	"common"
	"database/sql"
	"logs"
	//_ "mysql0"
	_"github.com/go-sql-driver/mysql"
	"sync"
)

var datasourceInstance DataSource
var lock sync.Mutex

type simpleDataSource struct {
	size int
}

var source []*sql.DB
var db_chan chan *sql.DB
var wait_chan chan bool

func GetDataSource(config string, size int) DataSource {
	if size < 1 {
		panic("连接池数量必须大于0")
	}
	lock.Lock()
	defer lock.Unlock()
	if datasourceInstance == nil {
		logs.Info("创建DataSource实例", true)
		datasourceInstance = &simpleDataSource{size}
	} else {
		return datasourceInstance
	}
	readConfig(config)
	logs.Info("设置DataSource连接池大小", size)
	db_chan = make(chan *sql.DB, size)
	source = make([]*sql.DB, size)
	wait_chan = make(chan bool, size)
	for i := 0; i < size; i++ {
		if db, ok := datasourceInstance.Connect(); ok {
			source[i] = db
			wait_chan <- true
		}
	}
	go func() {
		for {
			<-wait_chan
			if len(source) > 0 {
				db := source[0]
				source = append(source[1:])
				logs.Info("向通道内推送一个数据库连接", true)
				db_chan <- db
			}

		}
	}()
	return datasourceInstance
}

func (dataSource *simpleDataSource) GetConnection() *sql.DB {
	logs.Info("获取DataSource连接")
	return <-db_chan
}
func (dataSource *simpleDataSource) DesConnection(db *sql.DB) {
	logs.Info("释放DataSource连接")
	source = append(source, db)
	wait_chan <- true
}

//连接数据库
func (baseDao *simpleDataSource) Connect() (*sql.DB, bool) {
	db, err := sql.Open("mysql", common.Get("database"))
	if err == nil {
		//logs.Info("数据库连接成功!")
		return db, true
	} else {
		logs.Info("数据库连接失败!", err)
	}
	return nil, false
}
func readConfig(config string) {
	conf := common.NewConfig([]string{config})
	conf.LoadConfig()
	logs.Info("seting files loaded successfully", common.Get("database"))
}

//关闭连接
/*func (baseDao *BaseDao) CloseConn() {
	conn := connSource.GetConnection()
	defer connSource.DesConnection(conn)
	conn.Close()
	fmt.Println("数据库连接关闭")
}*/
