package configuration

import (
	"cx-20-api/logs"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// TODO : make IP struct

type Barco struct {
	Network struct {
		IP         string
		SubnetMask string
		Gateway    string
	}
	Personalisation struct {
		OnScreenID struct {
			Language           string
			MeetingRoomName    string
			Location           string
			WelcomeMessage     string
			ShowNetworkInfo    bool
			EnableThreaterMode bool
		}
		Wallpaper struct {
			Number int
		}
	}
	WifiNetwork struct {
		LanSettings struct {
			LanHostName struct {
				Hostname string
			}
			PrimaryInterface struct {
				Method                     string
				PrimaryInterfaceIP         string
				PrimaryInterfaceSubnetMask string
				PrimaryInterfaceGateway    string
				PrimaryInterfaceDnsServer  string
			}
		}
		Services struct {
			ShareViaAirPlay    bool
			ShareViaGoogleCast bool
		}
		WirelessNetwork struct {
			SsidName     string
			Wpa2Password string
		}
	}
}

func LoadConfiguration(file string) Barco {
	var config Barco
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Config file is successfully loaded !")

	return config
}
