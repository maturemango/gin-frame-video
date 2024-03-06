package conn

import (
	"fmt"
	"sync"
	"time"
	"gin-frame/build/utils"

	_ "github.com/go-sql-driver/mysql" //载入 driver mysql
	"xorm.io/xorm"
)

var (
	once sync.Once
	engine *xorm.Engine
)

func NewXormDB(host, database, username, password string) *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s", 
	username, password, host, database, true, "Local")

	engine, err :=  xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic("new xorm db failed:" + err.Error())
	}
	if err := engine.Ping(); err != nil {
		panic("xorm db ping failed:" + err.Error())
	}
	engine.TZLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		engine.TZLocation = time.FixedZone("CST", 8*3600)
	}
	engine.SetMaxOpenConns(300)
	return engine
}

func InitDBEngine() error {
	once.Do(func() {
		engine = initDBEngine()
	})
	engine.ShowSQL(utils.Config.Mysql.ShowSQL)
	return engine.Ping()
}

func initDBEngine() *xorm.Engine {
	engine = NewXormDB(utils.Config.Mysql.Host, 
		utils.Config.Mysql.DataBase, 
		utils.Config.Mysql.UserName, 
		utils.Config.Mysql.Password)
	
	return engine
}

func GetEngine() *xorm.Engine {
	if engine == nil {
		engine = initDBEngine()
	}
	return engine
}