package test

import (
	"fmt"
	"regexp"
	"testing"

	"gin-frame/build/cmd"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	_ "github.com/spf13/viper"
)

func init() {
	cmd.LoadConfig()
	utils.InitConfig()
}

func TestCreateToken(t *testing.T) {
	// viper.SetDefault("mysql.host", "127.0.0.1:3306")
	// viper.SetDefault("mysql.database", "gin")
	// viper.SetDefault("mysql.username", "root")
	// viper.SetDefault("mysql.password", "123456")

	// 17606162963 role_id:1
	// 19443502652 role_id:2
	// 17852369877 role_id:3
	
	var usr model.UserInfo
	if _, err := conn.GetEngine().Where("phone = ?", "17606162963").Get(&usr); err != nil {
		t.Logf("mysql err: %s", err)
	}
	token, err := handlers.CreateToken(usr)
	if err != nil {
		t.Logf("create token err: %s", err)
	}
	fmt.Printf("token is: %v\n", token)
}

// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJOYW1lIjoiIiwiRXhwcmllQXQiOjE3MDY5NDc0MjYsImlzcyI6Im51bGwiLCJleHAiOjE3MDY5NDc0MjZ9.YRzmkOFpr_sZTgZXqPUg10q--5xuVyjLlFNEXAecUA490UHjQjWSfOwHjg4UbwK1Eo7fMew6PshHpgrgQHtsfw

func TestVerfiyToken(t *testing.T) {
	claim, err := handlers.VerfiyToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJOYW1lIjoiIiwiRXhwcmllQXQiOjE3MDgzOTI5MzQsImlzcyI6Im51bGwiLCJleHAiOjE3MDgzOTI5MzR9.HaAACG9n7WymP7Agu_QedNXc2yTjv7y26-nqx9FkqHGbQfsfA9LEWErlRnkZapajiheiNEbanBuKtc9fHMNHxw")
	if err != nil {
		t.Logf("token err: %s", err)
	}
	fmt.Println("token valid")
	if len(claim.UserName) <= 0 {
		fmt.Printf("claim: %v\n", claim.UserID)
	} else {
		fmt.Printf("claim name is: %v\n", &claim.UserName)
	}
}

func TestMatchPhone(t *testing.T) {
	phone := "19345678912"
	ok, err := regexp.MatchString("^1[3-9]\\d{9}$", phone)
	if len(phone) != 11 || !ok || err != nil {
		t.Fatal(fmt.Errorf("invalid phone"))
	}

	code := utils.RandomLoginCode(6)
	fmt.Println("code:" + code)
}