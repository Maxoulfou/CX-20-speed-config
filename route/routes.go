package route

import (
	"cx-20-api/entity"
	"fmt"
)

const (
	SYSTEM_STATUS              = "SystemStatus"
	REBOOT                     = "Reboot"
	PERSONALIZATION            = "Personalization"
	GET_WALLPAPER              = "GetWallpaper"
	UPDATE_WALLPAPER           = "UpdateWallpaper"
	UPDATE_SERVICE_AIR_PLAY    = "UpdateServiceAirPlay"
	UPDATES_ERVICE_GOOGLE_CAST = "UpdateServiceGoogleCast"
	UPDATE_HOSTNAME            = "UpdateHostName"
)

func RoutesDictionary() map[string]entity.Route {
	RoutesDict := make(map[string]entity.Route)
	RoutesDict["SystemStatus"] = entity.Route{Path: "/configuration/system/status", Method: "GET"}
	RoutesDict["Reboot"] = entity.Route{Path: "/operations/reboot", Method: "POST"}
	RoutesDict["Personalization"] = entity.Route{Path: "/configuration/personalization", Method: "PATCH"}
	RoutesDict["GetWallpaper"] = entity.Route{Path: "/configuration/wallpapers", Method: "GET"}
	RoutesDict["UpdateWallpaper"] = entity.Route{Path: "/configuration/wallpapers/selected", Method: "PATCH"}
	RoutesDict["UpdateServiceAirPlay"] = entity.Route{Path: "/configuration/features/airplay", Method: "PATCH"}
	RoutesDict["UpdateServiceGoogleCast"] = entity.Route{Path: "/configuration/features/google-cast", Method: "PATCH"}
	RoutesDict["UpdateHostName"] = entity.Route{Path: "/configuration/system/network", Method: "PATCH"}
	fmt.Printf("RoutesDict: %+v\n", RoutesDict)

	return RoutesDict
}
