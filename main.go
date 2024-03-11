package main

import (
	"gin-frame/build/cmd"
	"gin-frame/webapi/model"
	"log"
	"os"
)

var (
	Version    string
	GoVersion  string
	GitHash    string
	BuildTime  string
)

func main() {
	model.Version = Version
	model.GoVersion = GoVersion
	model.GitHash = GitHash
	model.BuildTime = BuildTime
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)  // 打印并终止程序
	}

	if flag := cmd.RootCmd.Flags().Lookup("help"); flag != nil && flag.Changed == true {
		os.Exit(0)
	}
}