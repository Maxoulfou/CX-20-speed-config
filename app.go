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

func Init() {
	var YamlEnv entity.YmlConfig
	YamlEnv.GetConfig()
	var BarcoConfig entity.Barco
	BarcoConfig.GetConfig()

	fmt.Println(YamlEnv)

	logs.WriteLogs("info", "App is started\n", true)

	if YamlEnv.Env == "debug" {
		_, file, line, _ := runtime.Caller(1)
		// TODO : if in yaml file, env is set up to debug, not prod
		logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") TEST ERROR LOG", false)
		logs.WriteLogs("warn", "("+file+" : "+strconv.Itoa(line)+") TEST WARN LOG", false)
		logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") TEST INFO LOG", false)
		logs.WriteLogs("debug", "("+file+" : "+strconv.Itoa(line)+") TEST DEBUG LOG", false)
	}
	config, _ := format.PrettyStruct(BarcoConfig)

	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Loaded configuration:\n"+config+"\n", true)
	logs.WriteLogs("info", "--- End loaded configuration ---", false)

	if api.CheckIfBarcoCxApiIsReachable() {
		fmt.Println("it is ok")
	} else {
		fmt.Println("it is not ok")

		return
	}

	prettyCfg, _ := format.PrettyStruct(YamlEnv)
	logs.WriteLogs("info", "yml config: \n"+prettyCfg+"\n", true)

	// api.Reboot()
	api.Personalization()
	fmt.Println("Get Wallpaper List")
	api.GetWallpaperList()
	api.ChangeWallpaper()
	api.UpdateAirplayService()
	api.UpdateGoogleCastService()

	os.Exit(0)
}
