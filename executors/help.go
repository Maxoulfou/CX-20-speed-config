package executors

import "fmt"

func JumpLine() {
	fmt.Printf("\n")
}

func Help() {
	JumpLine()
	fmt.Println("CX-20 API Help command")
	fmt.Println("Availible commands:")
	fmt.Printf("\t- all : Execute all commands\n")
	fmt.Printf("\t- airplay : Update airplay settings\n")
	fmt.Printf("\t- change-wallpaper : Change Wallpaper\n")
	fmt.Printf("\t- google-cast : Update airplay settings\n")
	fmt.Printf("\t- hostname : Update Hostname\n")
	fmt.Printf("\t- personalization : Update all personalization settings\n")
	fmt.Printf("\t- reboot : Reboot\n")
	fmt.Printf("\t- wallpaper-upload : Upload wallpaper\n")
	fmt.Printf("\t- wifi : Update wifi settings\n")
}
