package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	to := time.Now().Add(30 * time.Second)
	conn.SetReadDeadline(to)
	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	var data []byte
	_, err := conn.Read(data)
	conn.Write(data)
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
	}
}

func main() {
	// log := Logger{""}
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		// log.Error("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		// log.Error("Error accepting connection: " + err.Error())
		os.Exit(1)
	}
	go handleConnection(conn)
}
