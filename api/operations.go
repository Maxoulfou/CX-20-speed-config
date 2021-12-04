package api

import (
	"github.com/maxoulfou/CX-20-speed-config/entity"
	"github.com/maxoulfou/CX-20-speed-config/logs"
	"github.com/maxoulfou/CX-20-speed-config/route"
	"strconv"
)

var Dict = route.RoutesDictionary(false) // false by default

func SystemInformation() string {
	GetSystemInfo, GetSystemInfoError := MakeRequest(Dict[route.SYSTEM_STATUS].Path, "", Dict[route.SYSTEM_STATUS].Method)
	if GetSystemInfoError != nil {
		logs.WriteLogs("error", "MakeRebootRequestError: "+GetSystemInfoError.Error(), true)

		return GetSystemInfoError.Error()
	}
	logs.WriteLogs("info", "GetSystemInfo Request: "+GetSystemInfo.Status, true)

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

// UploadWallpaper : accepted link : https://url.com/photo.jpg or https://url.com/photo.png
func UploadWallpaper() {
	var cfg entity.Barco
	cfg.GetConfig()
	UploadWallpaperRequest, UploadWallpaperRequestError := MakeRequest(Dict[route.UPLOAD_WALLPAPER].Path, `{"url": "`+cfg.Personalisation.Wallpaper.Link+`"}`, Dict[route.UPLOAD_WALLPAPER].Method)
	if UploadWallpaperRequestError != nil {
		logs.WriteLogs("error", UploadWallpaperRequestError.Error(), true)

		return
	}

	logs.WriteLogs("info", "UploadWallpaperRequest: "+UploadWallpaperRequest.Status, true)

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
			"passphrase": ""
		  }
		}

		{
		  "accessPoint": {
			"ssid": "ClickShare-NEXTVISION"
			"passphrase": ""
		  }
		}
	*/

	var cfg entity.Barco
	cfg.GetConfig()
	UpdateSsid, UpdateSsidError := MakeRequest(Dict[route.UPDATE_SSID].Path, `{"accessPoint": { "ssid": "`+cfg.WifiNetwork.WirelessNetwork.SsidName+`", "WPA2passphrase": "`+cfg.WifiNetwork.WirelessNetwork.Wpa2Password+`"}}`, Dict[route.UPDATE_SSID].Method)
	if UpdateSsidError != nil {
		logs.WriteLogs("error", UpdateSsidError.Error(), true)

		return
	}

	logs.WriteLogs("info", "UpdateSsid: "+UpdateSsid.Status, true)
	return
}
