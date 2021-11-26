package entity

// Config is struct for yml config env and different state or choice
// env : (debug,test,pre-prod,prod,alpha,beta)
// log : (all-in-one, separated)
type Config struct {
	Env         string `yaml:"env"`
	Log         string `yaml:"log"`
	ApiUrl      string `yaml:"api-url"`
	ApiUser     string `yaml:"api-user"`
	ApiPassword string `yaml:"api-password"`
	ApiToken    string `yaml:"api-token"`
}

// Barco is struct for json config. It will be used for configure your
// CX-20 clickshare module
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
