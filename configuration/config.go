package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO : make IP struct

type Barco struct {
	Network struct {
		IP        string
		Broadcast string
		Gateway   string
	}
	Info struct {
		Name           string
		WelcomeMessage string
		PieceName      string
	}
	WirelessNetwork struct {
		SSIDName     string
		WPA2Password string
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
	return config
}
