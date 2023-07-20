package main

import (
	"log"
	"net"
	"os"
)

const (
	StopCharacter = "\r\n\r\n"
)

func SendToSocket(addr string, message string) {
	conn, err := net.Dial("unix", addr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}
