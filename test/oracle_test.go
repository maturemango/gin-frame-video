package test

import (
	"fmt"
	"gin-frame/build/cmd"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"testing"
)

func TestInitORADB(t *testing.T) {
	cmd.LoadConfig()
	utils.InitConfig()
	
	if err := conn.InitORADB(); err != nil {
		t.Logf("init ora failed:%v", err)
	}
	fmt.Println("connect success")
}