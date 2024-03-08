package cmd

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	pwd, _ := os.Executable()
	viper.AddConfigPath(path.Join(filepath.Dir(pwd), "conf"))
	viper.AddConfigPath("D:\\gopath\\gin\\gin-frame\\bin\\conf")  // 根据不同设备及时更换目录
	viper.AddConfigPath("./conf")
	viper.SetConfigName("base")
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("The configuration file was modified")
	})
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}