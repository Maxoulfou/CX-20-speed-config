package api

import (
	"cx-20-api/entity"
	"cx-20-api/logs"
	"cx-20-api/route"
	"strconv"
)

func SystemInformation() string {
	GetSystemInfo, GetSystemInfoError := MakeRequest(route.SystemStatus, "", "GET")
	if GetSystemInfoError != nil {
		logs.WriteLogs("error", "MakeRebootRequestError: "+GetSystemInfoError.Error(), true)

		return GetSystemInfoError.Error()
	}
	logs.WriteLogs("info", "Make reboot: "+GetSystemInfo.Status, true)

	return GetSystemInfo.Status
}

// Reboot function will reboot Barco CX-20 station
func Reboot() {
	MakeRebootRequest, MakeRebootRequestError := MakeRequest(route.Reboot, "", "POST")
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
	RequestPersonalization, RequestPersonalizationError := MakeRequest(route.Personalization, `{"language": "`+cfg.Personalisation.OnScreenID.Language+`","location": "`+cfg.Personalisation.OnScreenID.Location+`","meetingRoomName": "`+cfg.Personalisation.OnScreenID.MeetingRoomName+`","screensaver": {"timeout": "`+cfg.Personalisation.OnScreenID.ScreenSaverTimeout+`"},"showNetworkInfo": `+strconv.FormatBool(cfg.Personalisation.OnScreenID.ShowNetworkInfo)+`,"theaterMode": {"enabled": `+strconv.FormatBool(cfg.Personalisation.OnScreenID.EnableThreaterMode)+`},"welcomeMessage": "`+cfg.Personalisation.OnScreenID.WelcomeMessage+`"}`, "PATCH")
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
	UpdateHostName, UpdateHostNameError := MakeRequest(route.UpdateHostName, `{"hostname": "`+cfg.WifiNetwork.LanSettings.LanHostName.Hostname+`"}`, "PATCH")
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
	GetWallpaperListRequest, GetWallpaperListRequestError := MakeRequest(route.GetWallpaper, "", "GET")
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
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(route.UpdateWallpaper, `{"id": "`+cfg.Personalisation.Wallpaper.Number+`"}`, "PATCH")
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
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(route.UpdateServiceAirPlay, `{"enabled": "`+strconv.FormatBool(cfg.WifiNetwork.Services.ShareViaAirPlay)+`"}`, "PATCH")
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
	ChangeWallpaper, ChangeWallpaperError := MakeRequest(route.UpdateServiceGoogleCast, `{"enabled": "`+strconv.FormatBool(cfg.WifiNetwork.Services.ShareViaGoogleCast)+`"}`, "PATCH")
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
