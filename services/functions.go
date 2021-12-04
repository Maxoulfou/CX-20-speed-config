package services

import (
	"fmt"
	"github.com/maxoulfou/CX-20-speed-config/api"
)

func All() {
	Personalization()
	UpdateHostname()
	UpdateSsid()
	Airplay()
	GoogleCast()
	Airplay()
	ChangeWallpaper()
	fmt.Printf("\n--> all command exit with no errors <--\n")
}

func Airplay() {
	api.UpdateAirplayService()
	fmt.Printf("\nairplay command exit with no errors\n")
}

func GoogleCast() {
	api.UpdateGoogleCastService()
	fmt.Printf("\ngoogle-cast command exit with no errors\n")
}

func Reboot() {
	api.Reboot()
	fmt.Printf("\nreboot command exit with no errors\n")
}

func Personalization() {
	api.Personalization()
	fmt.Printf("\npersonalization command exit with no errors\n")
}

func UpdateHostname() {
	api.UpdateHostName()
	fmt.Printf("\nhostname command exit with no errors\n")
}

func UpdateSsid() {
	api.UpdateWifiSettings()
	fmt.Printf("\nwifi command exit with no errors\n")
}

func WallpaperUpload() {
	api.UploadWallpaper()
	fmt.Printf("\nwallpaper-upload command exit with no errors\n")
}

func ChangeWallpaper() {
	api.ChangeWallpaper()
	fmt.Printf("\nchange-wallpaper command exit with no errors\n")
}
