package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	checkInvalidInput()

	if os.Args[1] == "help" {
		help()
	} else if os.Args[1] == "listen" {
		listen()
	} else if os.Args[1] == "emit" {
		emit()
	}
}

func checkInvalidInput() {
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
	} else if os.Args[1] == "listen" {
		_, err := strconv.ParseInt(os.Args[2], 0, 64)
		if err != nil {
			fmt.Println("Invalid Port")
			os.Exit(1)
		}
	}

}

func help() {
	fmt.Println("Help Menu")
	fmt.Println("help : Opens help menu")
	fmt.Println("emit <ip> : Tries to get the revershell in the IP")
	fmt.Println("listen <port> : Listen for any connection in the port")
	fmt.Println()
	fmt.Println("Examples : ")
	fmt.Println("./reverseshell listen 8080")
	fmt.Println("./reverseshell emit 192.168.100.244")
}
