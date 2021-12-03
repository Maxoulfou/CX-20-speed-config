package executors

import "cx-20-api/api"

func All() {
	Personalization()
	UpdateHostname()
	UpdateSsid()
	Airplay()
	GoogleCast()
	Airplay()
	ChangeWallpaper()
}

func Airplay() {
	api.UpdateAirplayService()
}

func GoogleCast() {
	api.UpdateGoogleCastService()
}

func Reboot() {
	api.Reboot()
}

func Personalization() {
	api.Personalization()
}

func UpdateHostname() {
	api.UpdateHostName()
}

func UpdateSsid() {
	api.UpdateWifiSettings()
}

func WallpaperUpload() {
	api.UploadWallpaper()
}

func ChangeWallpaper() {
	api.ChangeWallpaper()
}
