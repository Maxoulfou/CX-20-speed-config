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

func LoadConfiguration(file string) (Barco Barco, Error error) {
	config := Barco

	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(configFile *os.File) {
		DeferConfigFileError := configFile.Close()
		if DeferConfigFileError != nil {
			_, file, line, _ := runtime.Caller(1)
			logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") Defering config file failed : "+DeferConfigFileError.Error())
		}
	}(configFile)

	jsonParser := json.NewDecoder(configFile)

	JsonParserDecodeError := jsonParser.Decode(&config)
	if JsonParserDecodeError != nil {
		_, file, line, _ := runtime.Caller(1)
		logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") Decode JSON parser failed : "+JsonParserDecodeError.Error())
		return Barco, JsonParserDecodeError
	}

	_, file, line, _ := runtime.Caller(1)
	logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Config file is successfully loaded !")

	return config, nil
}
