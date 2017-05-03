package config

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/atopse/comm/log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/atopse/comm"
)

var (
	// AppConfig 配置
	AppConfig config.Configer
	// AppPath 运行目录
	AppPath string
	// ConfigPath 配置文件全路径
	ConfigPath string
	// RunMode 运行模式
	RunMode string
)

func init() {
	cfgName := "app.conf"
	//检查是否是test环境下
	if flag.Lookup("test.v") != nil {
		cfgName = "app.test.conf"
		RunMode = "test"
		beego.BConfig.RunMode = "test"
	}

	isExist := func(name string) bool {
		_, err := os.Stat(name)
		if err == nil {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Panicln(err)
		return true
	}
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Panicln(err)
	}
	p := os.Getenv("AppConfigPath")
	if p != "" && isExist(p) {
		ConfigPath = p
	} else {
		workPath, err := os.Getwd()
		if err != nil {
			log.Panicln(err)
		}
		ConfigPath = filepath.Join(workPath, cfgName)
	}
	log.Debug(ConfigPath)
	if !isExist(ConfigPath) {
		ConfigPath = filepath.Join(AppPath, cfgName)
		if !isExist(ConfigPath) {
			if path, err := comm.SearchFile(ConfigPath); err == nil {
				ConfigPath = path
			}
		}
	}
	if !isExist(ConfigPath) {
		log.Panicln("未找到可用的配置文件", ConfigPath)
	}

	err = beego.LoadAppConfig("ini", ConfigPath)
	if err != nil {
		log.Panicln(err)
	}
	log.Info("应用配置文件:", ConfigPath)
	AppConfig = beego.AppConfig

	if RunMode == "test" {
		beego.BConfig.RunMode = "test"
		log.GetLogger().SetLevel(logs.LevelNotice)
	}
}
