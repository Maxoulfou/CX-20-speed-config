package api

import (
	"cx-20-api/entity"
	"cx-20-api/logs"
	"cx-20-api/route"
	"strconv"
)

var Dict = route.RoutesDictionary()

func SystemInformation() string {
	GetSystemInfo, GetSystemInfoError := MakeRequest(Dict[route.SYSTEM_STATUS].Path, "", Dict[route.SYSTEM_STATUS].Method)
	if GetSystemInfoError != nil {
		logs.WriteLogs("error", "MakeRebootRequestError: "+GetSystemInfoError.Error(), true)

		return GetSystemInfoError.Error()
	}
	logs.WriteLogs("info", "Make reboot: "+GetSystemInfo.Status, true)

	return GetSystemInfo.Status
}

// Reboot function will reboot Barco CX-20 station
func Reboot() {
	MakeRebootRequest, MakeRebootRequestError := MakeRequest(Dict[route.REBOOT].Path, "", Dict[route.REBOOT].Method)
	if MakeRebootRequestError != nil {
		logs.WriteLogs("error", "MakeRebootRequestError: "+MakeRebootRequestError.Error(), true)

		return
	}
	logs.WriteLogs("info", "Make reboot: "+MakeRebootRequest.Status, true)

	return
}

func Personalization() {
	var cfg entity.Barco
	cfg.GetConfig()
	RequestPersonalization, RequestPersonalizationError := MakeRequest(Dict[route.PERSONALIZATION].Path, `{"language": "`+cfg.Personalisation.OnScreenID.Language+`","location": "`+cfg.Personalisation.OnScreenID.Location+`","meetingRoomName": "`+cfg.Personalisation.OnScreenID.MeetingRoomName+`","screensaver": {"timeout": "`+cfg.Personalisation.OnScreenID.ScreenSaverTimeout+`"},"showNetworkInfo": `+strconv.FormatBool(cfg.Personalisation.OnScreenID.ShowNetworkInfo)+`,"theaterMode": {"enabled": `+strconv.FormatBool(cfg.Personalisation.OnScreenID.EnableThreaterMode)+`},"welcomeMessage": "`+cfg.Personalisation.OnScreenID.WelcomeMessage+`"}`, Dict[route.PERSONALIZATION].Method)
	if RequestPersonalizationError != nil {
		logs.WriteLogs("error", RequestPersonalizationError.Error(), true)

		return
	}

	logs.WriteLogs("info", "RequestPersonalization: "+RequestPersonalization.Status, false)

	// It will directly update hostname
	UpdateHostName()

	return
}

func UpdateHostName() {
	var cfg entity.Barco
	cfg.GetConfig()
	UpdateHostName, UpdateHostNameError := MakeRequest(Dict[route.UPDATE_HOSTNAME].Path, `{"hostname": "`+cfg.WifiNetwork.LanSettings.LanHostName.Hostname+`"}`, Dict[route.UPDATE_HOSTNAME].Method)
	if UpdateHostNameError != nil {
		logs.WriteLogs("error", UpdateHostNameError.Error(), true)

		return
	}

	logs.WriteLogs("info", "RequestPersonalization: "+UpdateHostName.Status, false)

	return
}

func GetWallpaperList() {
	var cfg entity.Barco
	cfg.GetConfig()
	GetWallpaperListRequest, GetWallpaperListRequestError := MakeRequest(Dict[route.GET_WALLPAPER].Path, "", Dict[route.GET_WALLPAPER].Method)
	if GetWallpaperListRequestError != nil {
		logs.WriteLogs("error", GetWallpaperListRequestError.Error(), true)

		return
	}

	logs.WriteLogs("info", "GetWallpaperListRequest: "+GetWallpaperListRequest.Status, true)
	return
}

func ChangeWallpaper() {
	var cfg entity.Barco
	cfg.GetConfig()
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(Dict[route.UPDATE_WALLPAPER].Path, `{"id": "`+cfg.Personalisation.Wallpaper.Number+`"}`, Dict[route.UPDATE_WALLPAPER].Method)
	if ChangeWallpaperError != nil {
		logs.WriteLogs("error", ChangeWallpaperError.Error(), true)

		return
	}

	logs.WriteLogs("info", "GetWallpaperListRequest: "+ChangeWallpaper.Status, true)
	return
}

func UpdateAirplayService() {
	var cfg entity.Barco
	cfg.GetConfig()
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(Dict[route.UPDATE_SERVICE_AIR_PLAY].Path, `{"enabled": "`+strconv.FormatBool(cfg.WifiNetwork.Services.ShareViaAirPlay)+`"}`, Dict[route.UPDATE_SERVICE_AIR_PLAY].Method)
	if ChangeWallpaperError != nil {
		logs.WriteLogs("error", ChangeWallpaperError.Error(), true)

		return
	}

	logs.WriteLogs("info", "GetWallpaperListRequest: "+ChangeWallpaper.Status, true)
	return
}

func UpdateGoogleCastService() {
	var cfg entity.Barco
	cfg.GetConfig()
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(Dict[route.UPDATES_ERVICE_GOOGLE_CAST].Path, `{"enabled": "`+strconv.FormatBool(cfg.WifiNetwork.Services.ShareViaGoogleCast)+`"}`, Dict[route.UPDATES_ERVICE_GOOGLE_CAST].Method)
	if ChangeWallpaperError != nil {
		logs.WriteLogs("error", ChangeWallpaperError.Error(), true)

		return
	}

	logs.WriteLogs("info", "GetWallpaperListRequest: "+ChangeWallpaper.Status, true)
	return
}

func UpdateWifiSettings() {
	/*
		{
		  "addressing": "Static",
		  "operationMode": "AccessPoint",
		  "ipAddress": "192.168.2.1",
		  "subnetMask": "255.255.255.0",
		  "accessPoint": {
		    "broadcastSsid": true,
		    "channel": 36,
		    "frequencyBand": "5 GHz",
		    "ssid": "ClickShare-NEXTVISION"
		  }
		}
	*/
}
