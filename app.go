package main

import (
	"cx-20-api/api"
	"cx-20-api/entity"
	"cx-20-api/executors"
	"cx-20-api/format"
	"cx-20-api/logs"
	"cx-20-api/route"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"runtime"
	"strconv"
)

var YamlEnv entity.YmlConfig
var BarcoConfig entity.Barco

func Init() {
	logs.WriteLogs("info", "App is started\n", true)

	AsciiWelcome()

	InitConfig()
	InitLogs()

	// Prettying yml and json config
	InitPrettyYamlConfig()
	InitLoadConfiguration()
	InitApiChecking()

	// Main program TODO : make args treatment
	InitCliTool()
	// InitProg()

	os.Exit(0)
}

func InitCliTool() {
	var Arguments []string
	Arguments = os.Args[1:]

	if len(Arguments) == 0 {
		fmt.Println("There is no argument, please execute './cx-20-api.exe executors'\n")
	} else if len(Arguments) == 1 {
		switch Arguments[0] {
		case "help":
			executors.Help()
			break
		case "reboot":
			executors.Reboot()
			break
		case "personalization":
			executors.Personalization()
			break
		case "airplay":
			executors.Airplay()
			break
		case "google-cast":
			executors.GoogleCast()
			break
		case "wallpaper-upload":
			executors.WallpaperUpload()
			break
		case "change-wallpaper":
			executors.ChangeWallpaper()
			break
		case "hostname":
			executors.UpdateHostname()
			break
		case "wifi":
			executors.UpdateSsid()
			break
		case "all":
			executors.All()
			break
		default:
			fmt.Printf("\nPlease refer you to executors argument")
		}
	} else if len(Arguments) > 1 {
		fmt.Printf("There is too much arguments, please refer you to executors argument\n")
	}
}

// AsciiWelcome will displaying welcome message - ascii art
func AsciiWelcome() {
	myFigure := figure.NewFigure("CX-20  API", "", true)
	myFigure.Print()
	fmt.Printf("Version: %+v\n", entity.VERSION)
}

// InitApiChecking check if CX-20 station's API is reachable
func InitApiChecking() {
	if api.CheckIfBarcoCxApiIsReachable() {
		if YamlEnv.Env == "debug" {
			fmt.Print("it is ok")
		}

	} else {
		if YamlEnv.Env == "debug" {
			fmt.Print("it is not ok")
		}

		return
	}
}

// InitConfig will load env and CX-20 station's configuration
func InitConfig() {
	YamlEnv.GetConfig()
	BarcoConfig.GetConfig()
}

// InitLoadConfiguration prettify JSON configuration
func InitLoadConfiguration() {
	config, _ := format.PrettyStruct(BarcoConfig)

	// TODO : remake stacktrace export function
	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Loaded configuration:\n"+config, true)
	fmt.Printf("\n------------------------------------------------------------------------------")
	logs.WriteLogs("info", "--- End loaded configuration ---", false)
}

// InitLogs init test log function
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

// InitPrettyYamlConfig prettify yml configuration file
func InitPrettyYamlConfig() {
	prettyCfg, _ := format.PrettyStruct(YamlEnv)
	logs.WriteLogs("info", "yml config: \n"+prettyCfg+"", true)
}

// InitProg is main function - procedural
func InitProg() {
	api.Reboot()
	api.Personalization()
	fmt.Println("Get Wallpaper List")
	api.GetWallpaperList()
	api.ChangeWallpaper()
	api.UpdateAirplayService()
	api.UpdateGoogleCastService()
	api.UploadWallpaper()
	route.RoutesDictionary(true)
	api.UpdateWifiSettings()
}
