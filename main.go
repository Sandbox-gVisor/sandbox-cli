package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")
	port := parser.Int("p", "port", &argparse.Options{Required: true, Help: "Socket port"})
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Socket address"})

	fuckFlag := parser.Flag("f", "fuck", &argparse.Options{Required: false, Help: "Send fuck json"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *fuckFlag {
		message := Message{Data: "hi", Type: "log"}
		SendToSocker(*address, *port, message.ToString())
	}

	fmt.Println(*fuckFlag)
}
