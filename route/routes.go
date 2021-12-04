package route

import (
	"bytes"
	"fmt"
	"github.com/maxoulfou/CX-20-speed-config/entity"
	"github.com/maxoulfou/CX-20-speed-config/logs"
)

const (
	SYSTEM_STATUS              = "SystemStatus"
	REBOOT                     = "Reboot"
	PERSONALIZATION            = "Personalization"
	UPLOAD_WALLPAPER           = "UploadWallpaper"
	GET_WALLPAPER              = "GetWallpaper"
	UPDATE_WALLPAPER           = "UpdateWallpaper"
	UPDATE_SERVICE_AIR_PLAY    = "UpdateServiceAirPlay"
	UPDATES_ERVICE_GOOGLE_CAST = "UpdateServiceGoogleCast"
	UPDATE_SSID                = "UpdateSsid"
	UPDATE_HOSTNAME            = "UpdateHostName"
)

func RoutesDictionary(console bool) map[string]entity.Route {
	RoutesDict := make(map[string]entity.Route)
	RoutesDict["SystemStatus"] = entity.Route{Path: "/configuration/system/status", Method: "GET"}
	RoutesDict["Reboot"] = entity.Route{Path: "/operations/reboot", Method: "POST"}
	RoutesDict["Personalization"] = entity.Route{Path: "/configuration/personalization", Method: "PATCH"}
	RoutesDict["UploadWallpaper"] = entity.Route{Path: "/configuration/wallpapers", Method: "POST"}
	RoutesDict["GetWallpaper"] = entity.Route{Path: "/configuration/wallpapers", Method: "GET"}
	RoutesDict["UpdateWallpaper"] = entity.Route{Path: "/configuration/wallpapers/selected", Method: "PATCH"}
	RoutesDict["UpdateServiceAirPlay"] = entity.Route{Path: "/configuration/features/airplay", Method: "PATCH"}
	RoutesDict["UpdateServiceGoogleCast"] = entity.Route{Path: "/configuration/features/google-cast", Method: "PATCH"}
	RoutesDict["UpdateSsid"] = entity.Route{Path: "/configuration/system/network/wireless/1", Method: "PATCH"}
	RoutesDict["UpdateHostName"] = entity.Route{Path: "/configuration/system/network", Method: "PATCH"}

	if console {
		// keyValuePair will get all routes in dict : [key] = [value]
		keyValuePair := createKeyValuePairs(RoutesDict)
		logs.WriteLogs("info", "RoutesDict : "+keyValuePair, true)
	}

	return RoutesDict
}

func createKeyValuePairs(m map[string]entity.Route) string {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "\n")
	for key, value := range m {
		fmt.Fprintf(b, "\t˫")
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	fmt.Fprintf(b, "------------------------------------------------------------------------------")
	return b.String()
}
