package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Socket address"})
	requestType := parser.String("t", "type", &argparse.Options{Required: true, Help: "requestType"})

	fileFlag := parser.Flag("f", "file", &argparse.Options{Required: false, Help: "Read from file"})
	content := parser.String("c", "conf", &argparse.Options{Required: true, Help: "Callbacks. Read from file or from it "})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	byteContent := []byte(*content)

	if *fileFlag {
		byteContent = ReadFile(string(byteContent))
	}
	SendToSocket(*address, *requestType, ReadDtos(byteContent))
}
