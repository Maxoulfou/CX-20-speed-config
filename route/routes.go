package route

var SystemStatus = "/configuration/system/status"                   // GET
var Reboot = "/operations/reboot"                                   // POST
var Personalization = "/configuration/personalization"              // PATCH
var GetWallpaper = "/configuration/wallpapers"                      // GET
var UpdateWallpaper = "/configuration/wallpapers/selected"          // PATCH
var UpdateServiceAirPlay = "/configuration/features/airplay"        // PATCH
var UpdateServiceGoogleCast = "/configuration/features/google-cast" // PATCH
var UpdateHostName = "/configuration/system/network"                // PATCH
