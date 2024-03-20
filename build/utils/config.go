package utils

import (
	"github.com/spf13/viper"
)

var Config struct {
	Http  struct {
		Port int
	}   
	Mysql struct {
		Host      string
		DataBase  string
		UserName  string
		Password  string
		ShowSQL   bool
	}
	Redis struct {
		Addr      string
		UserName  string
		Password  string
		PoolSize  int
	}
	Oracle struct {
		Host         string
		Port         int
		Name         string
		Password     string
		ServiceName  string
	}
	Login struct {
		ExprieAt  int
		LoginKey  string
		No        string
	}
	Video struct {
		No    string
	}
	Casbin struct {
		ModelPath    string
		PolicyPath   string
	}
	Manage struct {
		RoleId    int
	}
}

func InitConfig() error {
	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}