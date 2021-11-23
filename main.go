package main

import (
	"cx-20-api/configuration"
	"cx-20-api/format"
	"fmt"
	"os"
)

func main() {
	// TODO : make directory to store config file, and assiciated func to get current path
	BarcoConfig := configuration.LoadConfiguration("config.json")
	config, _ := format.PrettyStruct(BarcoConfig)

	fmt.Print("Your config:\n")
	fmt.Println(BarcoConfig)
	fmt.Print("Your pretty config:\n")
	fmt.Println(config)
	os.Exit(0)
}
