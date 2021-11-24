package configuration

import (
	"cx-20-api/entity"
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfiguration(file string) (Barco entity.Barco, Error error) {
	config := Barco

	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(configFile *os.File) {
		DeferConfigFileError := configFile.Close()
		if DeferConfigFileError != nil {
			panic(DeferConfigFileError.Error())
			//_, file, line, _ := runtime.Caller(1)
			//logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") Defering config file failed : "+DeferConfigFileError.Error())
		}
	}(configFile)

	jsonParser := json.NewDecoder(configFile)

	JsonParserDecodeError := jsonParser.Decode(&config)
	if JsonParserDecodeError != nil {
		panic(JsonParserDecodeError.Error())
		//_, file, line, _ := runtime.Caller(1)
		//logs.WriteLogs("error", "("+file+" : "+strconv.Itoa(line)+") Decode JSON parser failed : "+JsonParserDecodeError.Error())
		return Barco, JsonParserDecodeError
	}

	//_, file, line, _ := runtime.Caller(1)
	//logs.WriteLogs("info", "("+file+" : "+strconv.Itoa(line)+") Config file is successfully loaded !")

	return config, nil
}
