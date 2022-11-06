package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"issuerserver/models/database"
	"issuerserver/pkg/db"
	"sync"
)

var once sync.Once
var _mysqlDb *gorm.DB

// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetMySqlDb() *gorm.DB {
	Init()
	return _mysqlDb
}

// pckage initer
func Init() {
	once.Do(func() {
		initMySql()
	})
}

func initMySql() {
	dsn := db.NewMySqlDb().GetDsn()
	var err error
	_mysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := _mysqlDb.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	//Init tables
	//_mysqlDb.AutoMigrate(&database.DidData{}, &database.ClaimData{}, &database.Offer{}, &database.Template{})
	_mysqlDb.AutoMigrate(&database.DidData{}, &database.ClaimData{})
}
