package api

import (
	"crypto/tls"
	"cx-20-api/configuration"
	"cx-20-api/format"
	"cx-20-api/logs"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MakeRequest is the generic function to execute CX20API requests
func MakeRequest(path string, params string, method string) (response *http.Response, error error) {
	cfg := configuration.GetEnv()

	// TODO : make insecure request trough http transport method only for knowing end-devices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	logs.WriteLogs("info", "MakeRequest start")
	url := cfg.ApiUrl + path

	// TODO : Add body to request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	Auth := "Basic " + cfg.ApiToken

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", Auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	status := res.Status

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	FinalBody := string(body)
	FinalBody, _ = format.PrettyString(FinalBody)

	logs.WriteLogs("info", "api-MakeRequest : \n"+FinalBody)

	switch status {
	case "200 OK":
		return response, nil
		break
	default:
		return nil, err
		break
	}

	/*CheckIfBarcoCxApiIsReachable()
	cfg := configuration.GetEnv()

	// TODO : make insecure request trough http transport method only for knowing end-devices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	Cx20ApiUrl := cfg.ApiUrl + path

	payload := strings.NewReader(params)

	client := &http.Client{}
	req, reqError := http.NewRequest(method, Cx20ApiUrl, payload)

	if reqError != nil {
		return nil, reqError
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("APIKEY", cfg.ApiToken) // add CX-20 api key

	response, responseError := client.Do(req)
	if responseError != nil {
		return nil, responseError
	}

	// defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return response, responseError
	}
	*/
	return response, error
}

// CheckIfBarcoCxApiIsReachable check if barco api is reachable with getting device identity.
func CheckIfBarcoCxApiIsReachable() bool {
	cfg := configuration.GetEnv()

	// TODO : make insecure request trough http transport method only for knowing end-devices
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	logs.WriteLogs("info", "CheckIfBarcoCxApiIsReachable start")
	url := cfg.ApiUrl + "/configuration/system/device-identity"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}

	Auth := "Basic " + cfg.ApiToken

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", Auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	status := res.Status

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	FinalBody := string(body)
	FinalBody, _ = format.PrettyString(FinalBody)

	logs.WriteLogs("info", "api-CheckIfBarcoCxApiIsReachable : \n"+FinalBody)

	switch status {
	case "200 OK":
		return true
		break
	default:
		return false
		break
	}
	return false
}
