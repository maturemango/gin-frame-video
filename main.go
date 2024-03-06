package main

import (
	"gin-frame/build/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)  // 打印并终止程序
	}

	if flag := cmd.RootCmd.Flags().Lookup("help"); flag != nil && flag.Changed == true {
		os.Exit(0)
	}
}