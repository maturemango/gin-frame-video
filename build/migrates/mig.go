package migrates

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	mig "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"xorm.io/xorm"
	xmigrate "xorm.io/xorm/migrate"
)

func ReadAllSQL() []string {
	var sqlstr []string
	sqlPath := "D:\\gopath\\gin\\gin-frame\\bin\\migrate\\doc"
	files, err := ioutil.ReadDir(sqlPath)
	if err != nil {
		return sqlstr
	}

	var sqlContent [][]byte
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			content, err := ioutil.ReadFile(filepath.Join(sqlPath, file.Name()))
			if err != nil {
				return sqlstr
			}
			sqlContent = append(sqlContent, content)
		}
	}
	for i, content := range sqlContent {
		if i == 0 {
			if strings.Contains(string(content), ";") {
				sqlstr = strings.Split(string(content), ";")
			} else {
				sqlstr = append(sqlstr, string(content))
			}
		} else {
			if strings.Contains(string(content), ";") {
				str := strings.Split(string(content), ";")
				sqlstr = append(sqlstr, str...)
			} else {
				sqlstr = append(sqlstr, string(content))
			}
		}
	}
	return sqlstr
}

var migrates = []*xmigrate.Migration{
	{
		ID: time.Now().Format("2006-01-02 15:04:05"),
		Migrate: func(sees *xorm.Engine) error {
			var err error
			var id int
			sqlStr := ReadAllSQL()
			for i, str := range sqlStr {
				if strings.EqualFold(str, "") {
					continue
				}
				_, err = sees.Exec(str)
				if err != nil {
					id = i
					break
				}
			}
			return fmt.Errorf("exec:%d failed:%v", id, err)
		},
		Rollback: func(sees *xorm.Engine) error {
			_, err := sees.Exec(`drop table copy_user`)
			return err
		},
	},
} // 不太适合用于迁移数据库

var options = &xmigrate.Options{
	TableName: "copy_migrate",
	IDColumnName: "timeId",
}

func XormMigrateData() {
	db, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		"root", "123456", "127.0.0.1:3306", "copy", true, "Local"))
	if err != nil {
		log.Fatalf("connect mysql failed:%v", err)
	}

	m := xmigrate.New(db, options, migrates)
	if err := m.Migrate(); err != nil {
		log.Fatalf("migrate failed:%v", err)
	}
	fmt.Println("migrate success")
}

var (
	migrateDir = flag.String("migrate.file", "./bin/migrate", "Migrate file directory ?")
	mysqlDsn   = flag.String("mysql.dsn", os.Getenv("MYSQL_DSN"), "Mysql dsn")
)

func MigateData() {
	flag.Parse()
	db, err := sql.Open("mysql", *mysqlDsn)
	if err != nil {
		log.Fatalf("connect mysql failed:%v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ping mysql failed:%v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("start mysql migrate failed:%v", err)
	}

	// pwd, _ := os.Executable()
	// fmt.Printf("pwd : %v", pwd)
	m, err := mig.NewWithDatabaseInstance(
		fmt.Sprintf("%v", *migrateDir), // 读取文件前缀加上file://(引入migrate中的file) 路径要对
		// 同时路径下的执行文件也要对 在windows生成migrate的执行文件需要用到工具，不会)
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("migrate failed:%v", err)
	}

	if err := m.Up(); err == nil || err == mig.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("data migtate success")
	os.Exit(0)
}
