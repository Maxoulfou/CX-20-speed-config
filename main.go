package main

import (
	"cx-20-api/configuration"
	"cx-20-api/format"
	"cx-20-api/logs"
	"fmt"
	"github.com/kataras/golog"
	"io"
	"os"
)

func main() {
	// TODO : make directory to store config file, and assiciated func to get current path
	BarcoConfig := configuration.LoadConfiguration("config.json")
	config, _ := format.PrettyStruct(BarcoConfig)

	// logs
	f := logs.NewLogFile()
	defer f.Close()

	golog.SetOutput(io.MultiWriter(f, os.Stdout))
	logs.WriteLogs("info", "\nLogs system successfully initialized\n")

	fmt.Print("Your config:\n")
	fmt.Println(BarcoConfig)
	fmt.Print("Your pretty config:\n")
	fmt.Println(config)
	os.Exit(0)
}
