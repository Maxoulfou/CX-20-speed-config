package api

import (
	"cx-20-api/configuration"
	"net/http"
	"strings"
	"time"
)

// MakeRequest is the generic function to execute CX20API requests
func MakeRequest(path string, params string, method string) (response *http.Response, error error) {
	CheckIfBarcoCxApiIsReachable()
	cfg := configuration.GetEnv()

	Cx20ApiUrl := cfg.ApiUrl + path

	payload := strings.NewReader(params)

	client := &http.Client{}
	req, reqError := http.NewRequest(method, Cx20ApiUrl, payload)

	if reqError != nil {
		return nil, reqError
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("APIKEY", cfg.ApiKey) // add CX-20 api key

	response, responseError := client.Do(req)
	if responseError != nil {
		return nil, responseError
	}

	// defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return response, responseError
	}

	return response, error
}

func CheckIfBarcoCxApiIsReachable() {
	cfg := configuration.GetEnv()
	BarcoApiUrl := cfg.ApiUrl
	URL := BarcoApiUrl + "/explorer/"
	timeout := 500 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	_, err := client.Get(URL)

	if err != nil {
		panic(err.Error())

		return
	}
}
