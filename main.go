package main

import (
	"cx-20-api/api"
	"cx-20-api/configuration"
	"cx-20-api/format"
	"cx-20-api/logs"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	envConfig := configuration.GetEnv()

	logs.WriteLogs("info", "App is started")

	if envConfig.Env == "debug" {
		_, file, line, _ := runtime.Caller(1)
		// TODO : if in yaml file, env is set up to debug, not prod
		logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") TEST ERROR LOG")
		logs.WriteLogs("warn", "("+file+" : "+strconv.Itoa(line)+") TEST WARN LOG")
		logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") TEST INFO LOG")
		logs.WriteLogs("debug", "("+file+" : "+strconv.Itoa(line)+") TEST DEBUG LOG")
		// Must be commented, this snippet broke code
		// logs.WriteLogs("fatal", "("+file+" : "+strconv.Itoa(line)+") TEST FATAL LOG")
	}

	BarcoConfig, _ := configuration.LoadConfiguration("config.json")
	config, _ := format.PrettyStruct(BarcoConfig)

	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Loaded configuration:\n"+config)
	logs.WriteLogs("info", "--- End loaded configuration ---")
	// Must be commented, this snippet broke code
	// test fatal log when I cant set loglevel :
	// logs.WriteLogs("", "fatal test")

	if api.CheckIfBarcoCxApiIsReachable() {
		fmt.Println("it is ok")
	} else {
		fmt.Println("it is not ok")
	}

	cfg := configuration.GetEnv()
	prettyCfg, _ := format.PrettyStruct(cfg)
	fmt.Printf("yml config: %+v\n", prettyCfg)

	fmt.Println("Test reboot")
	api.Reboot()

	os.Exit(0)
}
