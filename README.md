<h1 align="center">
  <br>
  <a href="https://www.barco.com/en/"><img src="https://www.cap-visio.com/wp-content/uploads/2020/11/Barco-Logo-900x244.png" alt="Logo Barco" width="200"></a>
  <br>
  CX-20-speed-config
  <br>
<a href="#"><img src="https://app.travis-ci.com/Maxoulfou/CX-20-speed-config.svg?branch=main" alt="Build status" width="100"></a>
</h1>

<h4 align="center">A tool to configure fastly Barco CX-20 station. It's free and open-source.</h4>

<hr>

# Table of Contents
1. [Configuration](#configuration)
   1. [config.json](#configjson)
   2. [env.yml](#envyml)
2. [Structures](#structures)
3. [Routes](#routes)
4. [Auth](#auth)


## Configuration

In this project, so far, I have decided to include two configuration files.
One that will be used to define the environments we are in. 

The other one will be used to define the basic information of the CX-20 station and the parameters to be applied to it, 
like for example, the welcome message, the localization or the language...

### config.json

> File to define the values to be injected into the station, as well as its basic information

The network section is only dedicated to the information for the good progress of the operations. There is no data in this file that will be applied to the station.

The sections just below are dedicated to the configuration to be applied to the station. Here they are:
- Personalization
  - OnScreenID
  - Wallpaper
- WifiNetwork
  - LanSettings
    - LanHostName
    - PrimaryInterface
  - Services
  - WirelessNetwork

### env.yml

> File to define the values of the current environment

File content :

```yml
env: "env-type"
log: "log-method"
```

`env:` valid options :
- debug
- test
- pre-prod
- prod
- alpha
- beta

`log:` valid options :
- all-in-one : all log level in same file
- separated : each log level separate by dedicated file

## Structures

> The structure below may not change, or there may be only additional possibilities

```go
type Config struct {
	Env string `yaml:"env"`
	Log string `yaml:"log"`
}
```

> The following structure may change in the future

```go
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
```

## Routes

> Not coded at this time

## Auth

> Not coded at this time

## LineCounter benchmark

```
/*
	BenchmarkBuffioScan   500      6408963 ns/op     4208 B/op    2 allocs/op
	BenchmarkBytesCount   500      4323397 ns/op     8200 B/op    1 allocs/op
	BenchmarkBytes32k     500      3650818 ns/op     65545 B/op   1 allocs/op
*/
```