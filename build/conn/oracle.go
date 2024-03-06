package conn

import (
	"database/sql"
	"fmt"
	"gin-frame/build/utils"

	_ "github.com/sijms/go-ora/v2"
	// go_ora "github.com/sijms/go-ora/v2"
)

var ora *sql.DB

func NewOracleDB(host, name, password, serviceName string, port int) *sql.DB {
	dsn := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", name, password, host, port,serviceName)
	// dsn := go_ora.BuildUrl(host, port, serviceName, name, password, nil)
	ora, err := sql.Open("oracle", dsn)
	if err != nil {
		panic("new oracle db failed:" + err.Error())
	}
	if err := ora.Ping(); err != nil {
		panic("oracle ping failed:" + err.Error())
	}

	ora.SetMaxOpenConns(300)
	return ora
}

func InitORADB() error {
	once.Do(func() {
		ora = initORADB()
	})

	if ora == nil {
		ora = initORADB()
	}

	return ora.Ping()
}

func initORADB() *sql.DB {
	return NewOracleDB(utils.Config.Oracle.Host,
		utils.Config.Oracle.Name,
		utils.Config.Oracle.Password,
		utils.Config.Oracle.ServiceName,
		utils.Config.Oracle.Port)
}
