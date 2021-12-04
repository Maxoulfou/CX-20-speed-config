package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/maxoulfou/CX-20-speed-config/api"
	"github.com/maxoulfou/CX-20-speed-config/entity"
	"github.com/maxoulfou/CX-20-speed-config/format"
	"github.com/maxoulfou/CX-20-speed-config/logs"
	"github.com/maxoulfou/CX-20-speed-config/route"
	"github.com/maxoulfou/CX-20-speed-config/services"
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
		fmt.Println("There is no argument, please execute './cx-20-api.exe help'\n")
	} else if len(Arguments) == 1 {
		switch Arguments[0] {
		case "help":
			services.Help()
			break
		case "reboot":
			services.Reboot()
			break
		case "personalization":
			services.Personalization()
			break
		case "airplay":
			services.Airplay()
			break
		case "google-cast":
			services.GoogleCast()
			break
		case "wallpaper-upload":
			services.WallpaperUpload()
			break
		case "change-wallpaper":
			services.ChangeWallpaper()
			break
		case "hostname":
			services.UpdateHostname()
			break
		case "wifi":
			services.UpdateSsid()
			break
		case "all":
			services.All()
			break
		default:
			fmt.Printf("\nPlease refer you to services argument")
		}
	} else if len(Arguments) > 1 {
		fmt.Printf("There is too much arguments, please refer you to services argument\n")
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
