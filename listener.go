package main

import (
	"fmt"
	"net"
	"os"
)

func listen() {
	dstream, err := net.Listen("tcp", ":"+os.Args[2])

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening at Port " + os.Args[2])
	defer dstream.Close()

	for {
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}
}

func handle(con net.Conn) {
	var command string
	for {
		fmt.Print("Command : ")
		fmt.Scan(&command)
		con.Write([]byte(command))
		buf := make([]byte, 1024)
		n, _ := con.Read(buf)
		data := string(buf[:n])
		fmt.Print(data)

	}
}
