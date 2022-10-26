package initialize

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"short-link-go/conf"
	"short-link-go/global"
)

func ViperInit() {
	var configFile string
	flag.StringVar(&configFile, "c", global.ConfigFile, "")
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic("配置文件解析失败！")
	}
	s := conf.ViperConfig{}
	uerr := v.Unmarshal(&s)
	if uerr != nil {
		panic("配置读取失败")
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		err := v.Unmarshal(&s)
		if err != nil {
			panic("配置重载失败")
		}
		global.CONF = s
	})
	global.CONF = s
}
