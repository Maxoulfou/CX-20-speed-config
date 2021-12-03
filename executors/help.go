package executors

import "fmt"

func JumpLine() {
	fmt.Printf("\n")
}

func Help() {
	JumpLine()
	fmt.Println("Help command run")
}
