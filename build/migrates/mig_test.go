package migrates

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
	mig "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"xorm.io/xorm"
	xmigrate "xorm.io/xorm/migrate"
)

func TestMigData(t *testing.T) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s", "root", "123456", "127.0.0.1:3306", "copy", true, "Local"))
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

	m, err := mig.NewWithDatabaseInstance(
		"file://D:/gopath/gin/gin-frame/bin/migrate",
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
}

func TestXormMigData(t *testing.T) {
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

var migrateions = []*xmigrate.Migration{
	{
		ID: "2024-03-15 13:58:01",
		Rollback: func(sees *xorm.Engine) error {
			_, err := sees.Exec("drop table copy_user")
			return err
		},
	},
}

func TestRollback(t *testing.T) {
	db, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		"root", "123456", "127.0.0.1:3306", "copy", true, "Local"))
	if err != nil {
		log.Fatalf("connect mysql failed:%v", err)
	}

	m := xmigrate.New(db, options, migrateions)
	if err := m.RollbackMigration(migrateions[0]); err != nil {
		log.Fatalf("roolback failed:%v", err)
	}
	fmt.Println("rollback success")
}

// func TestTime(t *testing.T) {
// 	ti := time.Now().Format("2006-01-02 15:04:05") // layout必须按照规定好的格式来
// 	fmt.Println(ti)
// }
