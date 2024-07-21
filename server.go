package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	logFactory "github.com/skepticmonk/gonetworking/logger"
)

func handleConnection(conn net.Conn, logger logFactory.Logger) {
	defer conn.Close()
	to := time.Now().Add(1 * time.Second)
	conn.SetDeadline(to)
	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	received := make([]byte, 4096)
	println("Reading data...")
	temp := make([]byte, 256)
	for {
		_, err := conn.Read(temp)
		logger.Info("Data:", string(temp), len(temp))
		if err != nil {
			if err == io.EOF {
				logger.Info("EOF")
				break
			}
			logger.Info("Read data failed:", err.Error())
			break
		}
		received = append(received, temp...)
	}
	logger.Info(string(received))
	if len(received) > 0 {
		_, err := conn.Write(received)
		if err != nil {
			logger.Error(err.Error())
		}
	}

}

func main() {
	log := logFactory.Logger{Name: ""}
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Error("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	for {
		log.Info("Ready to receive connection")
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, log)
	}
}
