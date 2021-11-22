package main

import "cx-20-api/configuration"

func main() {
	// TODO : make directory to store config file, and assiciated func to get current path
	configuration.LoadConfiguration("config.json")
}
