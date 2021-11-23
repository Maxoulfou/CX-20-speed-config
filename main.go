package main

import (
	"cx-20-api/configuration"
	"cx-20-api/format"
	"cx-20-api/logs"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// TODO : make directory to store config file, and assiciated func to get current path
	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") TEST ERROR LOG")
	logs.WriteLogs("warn", "("+file+" : "+strconv.Itoa(line)+") TEST WARN LOG")
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") TEST INFO LOG")
	logs.WriteLogs("debug", "("+file+" : "+strconv.Itoa(line)+") TEST DEBUG LOG")

	BarcoConfig := configuration.LoadConfiguration("config.json")
	config, _ := format.PrettyStruct(BarcoConfig)

	fmt.Print("Your config:\n")
	fmt.Println(BarcoConfig)
	fmt.Print("Your pretty config:\n")
	fmt.Println(config)
	os.Exit(0)
}
