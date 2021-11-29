package entity

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigIT interface {
	GetConfig()
}

func (receiver *YmlConfig) GetConfig() *YmlConfig {
	yamlFile, ErrReadYamlFile := ioutil.ReadFile("env-prod.yml")
	if ErrReadYamlFile != nil {
		log.Fatalf("YAML config.GetConfig ErrReadYamlFile: " + ErrReadYamlFile.Error())
	}

	ErrUnmarshalYamlFile := yaml.Unmarshal(yamlFile, receiver)
	if ErrUnmarshalYamlFile != nil {
		log.Fatalf("YAML config.GetConfig ErrUnmarshalYamlFile: " + ErrUnmarshalYamlFile.Error())
	}

	return receiver
}

type BarcoConfigIT interface {
	GetConfig()
}

func (receiver *Barco) GetConfig() *Barco {
	jsonFile, ErrReadJsonBarcoFile := ioutil.ReadFile("config.json")
	if ErrReadJsonBarcoFile != nil {
		log.Fatalf("JSON config.GetConfig ErrReadJsonBarcoFile: " + ErrReadJsonBarcoFile.Error())
	}

	ErrUnmarshalJsonBarcoFile := json.Unmarshal(jsonFile, receiver)
	if ErrUnmarshalJsonBarcoFile != nil {
		log.Fatalf("JSON config.GetConfig ErrUnmarshalJsonBarcoFile: " + ErrUnmarshalJsonBarcoFile.Error())
	}

	return receiver
}

// YmlConfig is struct for yml config env and different state or choice
type YmlConfig struct {
	Env         string `yaml:"env"`
	Log         string `yaml:"log"`
	ApiUrl      string `yaml:"api-url"`
	ApiUser     string `yaml:"api-user"`
	ApiPassword string `yaml:"api-password"`
	ApiToken    string `yaml:"api-token"`
}

// Barco is struct for json config. It will be used for configure your / CX-20 clickshare module
type Barco struct {
	// TODO : make ip struct
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
