package api

import (
	"cx-20-api/route"
	"fmt"
)

func Reboot() {
	request, err := MakeRequest(route.Reboot, "", "POST")
	if err != nil {
		return
	}
	fmt.Printf("MakeRequest reboot : %+v\n", request)
}
