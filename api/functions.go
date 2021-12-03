package api

import (
	"crypto/tls"
	"cx-20-api/entity"
	"cx-20-api/logs"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var YamlEnv entity.YmlConfig

// MakeRequest is the generic function to execute CX20API requests
func MakeRequest(path string, params string, method string) (response *http.Response, error error) {
	var cfg entity.Barco
	cfg.GetConfig()
	var EnvCfg entity.YmlConfig
	EnvCfg.GetConfig()

	// WARN : make insecure request trough http transport method only for knowing end-devices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := EnvCfg.ApiUrl + path

	payload := strings.NewReader(params)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

		return
	}

	Auth := "Basic " + EnvCfg.ApiToken
	req.Header.Add("Authorization", Auth)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

		return
	}

	logs.WriteLogs("info", string(body), true)

	return res, err
}

// CheckIfBarcoCxApiIsReachable TODO : remake this func prettiest as possible
// CheckIfBarcoCxApiIsReachable check if barco api is reachable with getting device identity.
func CheckIfBarcoCxApiIsReachable() bool {
	YamlEnv.GetConfig() // Load Yaml configuration
	System := SystemInformation()

	if YamlEnv.Env == "debug" {
		fmt.Printf("\nSystemStatus: %+v\n", System)
	}

	if System == "200 OK" {

		return true
	} else {

		return false
	}
}
