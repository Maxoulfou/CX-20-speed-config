package main

import (
	"cx-20-api/api"
	"cx-20-api/entity"
	"cx-20-api/format"
	"cx-20-api/logs"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var YamlEnv entity.YmlConfig
var BarcoConfig entity.Barco

func Init() {
	logs.WriteLogs("info", "App is started\n", true)

	InitConfig()
	InitLogs()
	InitLoadConfiguration()
	InitApiChecking()
	InitPrettyYamlConfig()
	InitProg()

	os.Exit(0)
}

func InitApiChecking() {
	if api.CheckIfBarcoCxApiIsReachable() {
		fmt.Println("it is ok")
	} else {
		fmt.Println("it is not ok")

		return
	}
}

func InitConfig() {
	YamlEnv.GetConfig()
	BarcoConfig.GetConfig()
}

func InitLoadConfiguration() {
	config, _ := format.PrettyStruct(BarcoConfig)

	// TODO : remake stacktrace export function
	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Loaded configuration:\n"+config+"\n", true)
	logs.WriteLogs("info", "--- End loaded configuration ---", false)
}

func InitLogs() {
	if YamlEnv.Env == "debug" {
		_, file, line, _ := runtime.Caller(1)
		// TODO : if in yaml file, env is set up to debug, not prod
		logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") TEST ERROR LOG", false)
		logs.WriteLogs("warn", "("+file+" : "+strconv.Itoa(line)+") TEST WARN LOG", false)
		logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") TEST INFO LOG", false)
		logs.WriteLogs("debug", "("+file+" : "+strconv.Itoa(line)+") TEST DEBUG LOG", false)
	}

}

func InitPrettyYamlConfig() {
	prettyCfg, _ := format.PrettyStruct(YamlEnv)
	logs.WriteLogs("info", "yml config: \n"+prettyCfg+"\n", true)
}

func InitProg() {
	// api.Reboot()
	api.Personalization()
	fmt.Println("Get Wallpaper List")
	api.GetWallpaperList()
	api.ChangeWallpaper()
	api.UpdateAirplayService()
	api.UpdateGoogleCastService()
}
