package main

import (
	"fmt"
	"os"
)

func main() {
	checkInput()
	fmt.Println("App started")
}

func checkInput() {
	if len(os.Args) == 1 {
		fmt.Println("Not Enough Arguments")
		os.Exit(1)
		return
	} else if len(os.Args) == 2 && os.Args[1] != "help" {
		fmt.Println("Invalid arguments")
		os.Exit(1)
		return
	} else if os.Args[1] != "help" && os.Args[1] != "listen" && os.Args[1] != "emit" {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	} else if (os.Args[1] == "listen" || os.Args[1] == "emit") && len(os.Args) == 2 {
		fmt.Println("Not Enough Arguments")
		os.Exit(1)
	}

}
